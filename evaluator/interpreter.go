package evaluator

import (
	"os/exec"
    "bytes"
	"strings"
    "fmt"

    "proxima/config"
)

var templates = map[ProgrammingLanguage]string{
    PYTHON: "%s\nprint(%s(%s))\n",
    JAVASCRIPT: "%s\nconsole.log(%s(%s));\n",
    LUA: "%s\nprint(%s(%s))\n",
    RUBY: "%s\nputs %s(%s)\n",
}

var unnamedArgTemplates = map[ProgrammingLanguage]string{
    PYTHON: "r'%s'",
    JAVASCRIPT: "'%s'",
    LUA: "'%s'",
    RUBY: "'%s'",
}

var namedArgTemplates = map[ProgrammingLanguage]string{
    PYTHON: "%s=r'%s'",
    JAVASCRIPT: "%s: '%s'",
    LUA: "%s = '%s'",
    RUBY: "%s: '%s'",
}

type Interpreter struct {
    languageToCommand map[ProgrammingLanguage]string
}

func NewInterpreter(config *config.Config) (*Interpreter, error) {
    languageToCommand := map[ProgrammingLanguage]string {
        PYTHON: config.Evaluator.PythonCmd,
        JAVASCRIPT: config.Evaluator.JavaScriptCmd,
        LUA: config.Evaluator.LuaCmd,
        RUBY: config.Evaluator.RubyCmd,
    }

    return &Interpreter{
        languageToCommand: languageToCommand,
    }, nil
}

func (i *Interpreter) Evaluate(args []struct{Name string; Value string}, component Component) (string, error) {
    formattedArgs := formatArgs(component.language, args)
    script := fmt.Sprintf(templates[component.language], component.content, component.name, formattedArgs)

    // execute script
    command := strings.Split(i.languageToCommand[component.language], " ")

    first := command[0]
    rest := command[1:]
    rest = append(rest, script)

    cmd := exec.Command(first, rest...)

    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf(stderr.String())
    }

    return strings.TrimSuffix(out.String(), "\n"), nil
}

func (i *Interpreter) Close() {
    // do nothing
}

func formatArgs(language ProgrammingLanguage, args []struct{Name string; Value string}) string {
    formattedArgs := make([]string, 0, len(args))
    for _, arg := range args {
        if arg.Value == "" {
            continue
        }

        value := strings.ReplaceAll(arg.Value, "'", SINGLE_QUOTE)
        value = strings.ReplaceAll(value, "\n", LINEBREAK)
        if strings.HasPrefix(arg.Name, "_unnamed_") {
            formattedArgs = append(formattedArgs, fmt.Sprintf(unnamedArgTemplates[language], value))
        } else {
            formattedArgs = append(formattedArgs, fmt.Sprintf(namedArgTemplates[language], arg.Name, value))
        }
    }

    return strings.Join(formattedArgs, ", ")
}
