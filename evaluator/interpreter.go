package evaluator

import (
    "os/exec"
    "bufio"
    "strings"
    "fmt"
    "io"
    "os"

    "proxima/config"

    _ "embed"
)

//go:embed persistent_interpreter.py
var embeddedPythonScript string

var interpreterToCommand = map[string][]string{
    "python": {"python", "-u"},
    "python3": {"python3", "-u"},
    "node": {"node", "-e"},
    "lua": {"lua", "-e"},
    "ruby": {"ruby", "-e"},
}


var templates = map[ProgrammingLanguage]string{
    PYTHON: "%s\nresult = %s(%s)\nprint('<<<START_RESULT>>>')\nprint(result)\nprint('<<<END_RESULT>>>')\n",
}

var unnamedArgTemplates = map[ProgrammingLanguage]string{
    PYTHON: "r'''%s'''",
}

var namedArgTemplates = map[ProgrammingLanguage]string{
    PYTHON: "%s=r'''%s'''",
}

type Interpreter struct {
    cmd               *exec.Cmd
    stdin             io.WriteCloser
    stdout            io.ReadCloser
    stderr            io.ReadCloser
    scriptPath        string
}

func createTempPythonScript(scriptContent string) (string, error) {
    tmpFile, err := os.CreateTemp("", "persistent_interpreter_*.py")
    if err != nil {
        return "", err
    }
    defer tmpFile.Close()

    _, err = tmpFile.Write([]byte(scriptContent))
    if err != nil {
        return "", err
    }

    return tmpFile.Name(), nil
}

func NewInterpreter(config *config.Config) (*Interpreter, error) {
    scriptPath, err := createTempPythonScript(embeddedPythonScript)
    if err != nil {
        return nil, err
    }

    fullCmd := interpreterToCommand[config.Runtimes.Python]
    cmd := exec.Command(fullCmd[0], append(fullCmd[1:], scriptPath)...)
    stdin, err := cmd.StdinPipe()
    if err != nil {
        return nil, err
    }
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        return nil, err
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        return nil, err
    }

    err = cmd.Start()
    if err != nil {
        return nil, err
    }

    return &Interpreter{
        cmd:               cmd,
        stdin:             stdin,
        stdout:            stdout,
        stderr:            stderr,
        scriptPath:        scriptPath,
    }, nil
}

func (i *Interpreter) Evaluate(args []struct{ Name string; Value string }, component Component) (string, error) {
    formattedArgs := formatArgs(component.language, args)
    script := fmt.Sprintf(templates[component.language], component.content, component.name, formattedArgs)

    stdin := i.stdin
    stdout := i.stdout
    stderr := i.stderr

    codeToSend := script + "\n<<<END>>>\n"
    _, err := io.WriteString(stdin, codeToSend)
    if err != nil {
        return "", err
    }

    reader := bufio.NewReader(stdout)
    var output strings.Builder
    var inResult bool

    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            // Read from stderr
            errOutput, _ := io.ReadAll(stderr)
            return "", fmt.Errorf("Error: %s", strings.TrimSpace(string(errOutput)))
        }

        line = strings.TrimSpace(line)

        if line == "<<<START_RESULT>>>" {
            inResult = true
            continue
        }
        if line == "<<<END_RESULT>>>" {
            break
        }
        if line == "<<<EXCEPTION>>>" {
            var exceptionOutput strings.Builder
            for {
                exceptionLine, err := reader.ReadString('\n')
                if err != nil {
                    break
                }
                exceptionLine = strings.TrimSpace(exceptionLine)
                if exceptionLine == "<<<END_EXCEPTION>>>" {
                    break
                }
                exceptionOutput.WriteString(exceptionLine + "\n")
            }
            return "", fmt.Errorf(exceptionOutput.String())
        }
        if inResult {
            output.WriteString(line)
            output.WriteString("\n")
        }
    }

    return strings.TrimSpace(output.String()), nil
}

func (i *Interpreter) Close() {
    if i.stdin != nil {
        i.stdin.Close()
    }
    if i.stdout != nil {
        i.stdout.Close()
    }
    if i.stderr != nil {
        i.stderr.Close()
    }
    if i.cmd != nil {
        i.cmd.Process.Kill()
        i.cmd.Wait()
    }
    if i.scriptPath != "" {
        os.Remove(i.scriptPath)
    }
}

func formatArgs(language ProgrammingLanguage, args []struct{ Name string; Value string }) string {
    formattedArgs := make([]string, 0, len(args))
    for _, arg := range args {
        if arg.Value == "" {
            continue
        }

        value := strings.ReplaceAll(arg.Value, "'''", "\\'''")
        value = strings.ReplaceAll(value, "\n", LINEBREAK)
        if strings.HasPrefix(arg.Name, "_unnamed_") {
            formattedArgs = append(formattedArgs, fmt.Sprintf(unnamedArgTemplates[language], value))
        } else {
            formattedArgs = append(formattedArgs, fmt.Sprintf(namedArgTemplates[language], arg.Name, value))
        }
    }

    return strings.Join(formattedArgs, ", ")
}
