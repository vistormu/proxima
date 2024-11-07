package evaluator

import (
	"os/exec"
    "bytes"
	"strings"
    "fmt"
    "sort"

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
    JAVASCRIPT: "%s='%s'",
    LUA: "%s='%s'",
    RUBY: "%s='%s'",
}

type Interpreter struct {
    languageToCommand map[ProgrammingLanguage]string
}

func NewInterpreter(config *config.Config) Interpreter {
    languageToCommand := map[ProgrammingLanguage]string {
        PYTHON: config.Evaluator.PythonCmd,
        JAVASCRIPT: config.Evaluator.JavaScriptCmd,
        LUA: config.Evaluator.LuaCmd,
        RUBY: config.Evaluator.RubyCmd,
    }

    return Interpreter{
        languageToCommand: languageToCommand,
    }
}

func (i *Interpreter) Evaluate(args map[string]string, component Component) (string, error) {
    // format args
    formattedArgs := formatArgs(component.language, args)

    // get script
    script := fmt.Sprintf(templates[component.language], component.content, component.name, formattedArgs)

    // execute script
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

    // Execute the command
    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf(stderr.String())
    }

    // remove the trailing newline given by the prints
    return strings.TrimSuffix(out.String(), "\n"), nil
}

func formatArgs(language ProgrammingLanguage, args map[string]string) string {
    // Convert map to a slice of anonymous structs to maintain order after sorting
    sortedArgs := make([]struct {
        Key   string
        Value string
    }, 0, len(args))

    for key, value := range args {
        sortedArgs = append(sortedArgs, struct {
            Key   string
            Value string
        }{Key: key, Value: value})
    }

    // Sort the slice by Key
    sort.Slice(sortedArgs, func(i, j int) bool {
        return sortedArgs[i].Key < sortedArgs[j].Key
    })

    // Initialize formattedArgs slice for formatted strings
    formattedArgs := make([]string, 0, len(args))

    // Iterate over the sorted slice
    for _, arg := range sortedArgs {
        name := arg.Key
        value := arg.Value
        if value == "" {
            continue
        }

        value = strings.ReplaceAll(value, "'", "\\'")
        if strings.HasPrefix(name, "_unnamed_") {
            formattedArgs = append(formattedArgs, fmt.Sprintf(unnamedArgTemplates[language], value))
        } else {
            formattedArgs = append(formattedArgs, fmt.Sprintf(namedArgTemplates[language], name, value))
        }
    }

    // Join the formatted arguments into a single string
    return strings.Join(formattedArgs, ", ")
}
