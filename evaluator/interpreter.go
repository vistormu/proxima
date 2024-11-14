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

var template = `%s
result = %s(%s)
print('<<<START_RESULT>>>')
print(result)
print('<<<END_RESULT>>>')
`

var unnamedArgTemplate = "r'''%s'''"
var namedArgTemplate = "%s=r'''%s'''"

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

    _, err = tmpFile.WriteString(scriptContent)
    return tmpFile.Name(), err
}

func NewInterpreter(config *config.Config) (*Interpreter, error) {
    scriptPath, err := createTempPythonScript(embeddedPythonScript)
    if err != nil {
        return nil, err
    }

    cmd := exec.Command(config.Evaluator.Python, "-u", scriptPath)
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

func (i *Interpreter) Evaluate(args []struct{ Name, Value string }, component Component) (string, error) {
    formattedArgs := formatArgs(args)
    script := fmt.Sprintf(template, component.content, component.name, formattedArgs)

    if _, err := io.WriteString(i.stdin, script+"\n<<<END>>>\n"); err != nil {
        return "", err
    }

    return i.readOutput()
}


func (i *Interpreter) readOutput() (string, error) {
    var output strings.Builder
    var inResult bool

    reader := bufio.NewReader(i.stdout)
    for {
        line, err := reader.ReadString('\n')
        if err != nil {
            errOutput, _ := io.ReadAll(i.stderr)
            return "", fmt.Errorf("Error: %s", strings.TrimSpace(string(errOutput)))
        }

        line = strings.TrimSpace(line)
        switch line {
        case "<<<START_RESULT>>>":
            inResult = true
        case "<<<END_RESULT>>>":
            return strings.TrimSpace(output.String()), nil
        case "<<<EXCEPTION>>>":
            return i.handleException(reader)
        default:
            if inResult {
                output.WriteString(line + "\n")
            }
        }
    }
}

func (i *Interpreter) handleException(reader *bufio.Reader) (string, error) {
    var exceptionOutput strings.Builder
    for {
        line, err := reader.ReadString('\n')
        if err != nil || strings.TrimSpace(line) == "<<<END_EXCEPTION>>>" {
            break
        }
        exceptionOutput.WriteString(strings.TrimSpace(line) + "\n")
    }
    return "", fmt.Errorf(exceptionOutput.String())
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

func formatArgs(args []struct{ Name, Value string }) string {
    var formattedArgs []string
    for _, arg := range args {
        if arg.Value == "" {
            continue
        }

        sanitizedValue := strings.ReplaceAll(strings.ReplaceAll(arg.Value, "'''", "\\'''"), "\n", LINEBREAK)
        if strings.HasPrefix(arg.Name, "_unnamed_") {
            formattedArgs = append(formattedArgs, fmt.Sprintf(unnamedArgTemplate, sanitizedValue))
        } else {
            formattedArgs = append(formattedArgs, fmt.Sprintf(namedArgTemplate, arg.Name, sanitizedValue))
        }
    }
    return strings.Join(formattedArgs, ", ")
}
