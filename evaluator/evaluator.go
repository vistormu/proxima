package evaluator

import (
	"fmt"
    "strings"
    "bytes"
    "os/exec"
    "sort"

    "proxima/config"
    "proxima/ast"
    "proxima/errors"
)

type Evaluator struct {
    expressions []ast.Expression
    components map[string]Component

    file string
    currentLine int

    evalCommands map[ProgrammingLanguage][]string
}

// PUBLIC
func New(expressions []ast.Expression, file string, config *config.Config) (*Evaluator, error) {
    // load components
    components, err := loadComponents(expressions, config)
    if err != nil {
        return nil, err
    }

    evalCommands := map[ProgrammingLanguage][]string{
        PYTHON: strings.Split(*config.Evaluator.PythonCmd, " "),
        JAVASCRIPT: strings.Split(*config.Evaluator.JavaScriptCmd, " "),
        LUA: strings.Split(*config.Evaluator.LuaCmd, " "),
        RUBY: strings.Split(*config.Evaluator.RubyCmd, " "),
    }

    return &Evaluator{
        expressions,
        components,
        file,
        0,
        evalCommands,
    }, nil
}

func (e *Evaluator) Evaluate() (string, error) {
    content := ""
    for _, expression := range e.expressions {
        result, err := e.evaluateExpression(expression)
        if err != nil {
            return "", err
        }
        content += result
    }

    return content, nil
}

// EVALUATION
func (e *Evaluator) evaluateExpression(expression ast.Expression) (string, error) {
    e.currentLine = expression.Line()
    switch expression := expression.(type) {
    case *ast.Text:
        return expression.Value, nil
    case *ast.Tag:
        return e.evaluateTag(expression)
    default:
        return "", nil
    }
}

func (e *Evaluator) evaluateTag(tag *ast.Tag) (string, error) {
    component := e.components[tag.Name]

    // evaluate args
    args := make(map[string]string, len(tag.Args))
    for i, arg := range tag.Args {
        evaluatedArg := ""
        for _, value := range arg.Values {
            result, err := e.evaluateExpression(value)
            if err != nil {
                return "", err
            }
            evaluatedArg += result
        }
        name := arg.Name
        if name == "" {
            name = fmt.Sprintf("_unnamed_%d", i)
        }
        args[name] = evaluatedArg
    }

    // format args
    formattedArgs := formatArgs(component.language, args)

    // get script
    script := getScript(component.language, component, formattedArgs)

    // execute script
    command := e.evalCommands[component.language]

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
        return "", errors.NewEvalError(errors.ERROR_EXECUTING_SCRIPT, e.file, e.currentLine, component.name, stderr.String())
    }

    // remove the trailing newline given by the prints
    return strings.TrimSuffix(out.String(), "\n"), nil
}

func getScript(language ProgrammingLanguage, component Component, args string) string {
    switch language {
    case PYTHON:
        return fmt.Sprintf("%s\nprint(%s(%s))", component.content, component.name, args)
    case JAVASCRIPT:
        return fmt.Sprintf("%s\nconsole.log(%s({%s}));", component.content, component.name, args)
    case LUA:
        return fmt.Sprintf("%s\nprint(%s(%s))", component.content, component.name, args)
    case RUBY:
        return fmt.Sprintf("%s\nputs %s(%s)", component.content, component.name, args)
    }

    return ""
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
            // add raw string to the formatted arguments // TODO: check if this is correct
            switch language {
            case PYTHON:
                formattedArgs = append(formattedArgs, fmt.Sprintf("r'%s'", value))
            case LUA:
                formattedArgs = append(formattedArgs, fmt.Sprintf("'%s'", value))
            case JAVASCRIPT:
                formattedArgs = append(formattedArgs, fmt.Sprintf("'%s'", value))
            case RUBY:
                formattedArgs = append(formattedArgs, fmt.Sprintf("'%s'", value))
            }
            continue
        }

        switch language {
        case PYTHON:
            formattedArgs = append(formattedArgs, fmt.Sprintf("%s=r'%s'", name, value))
        case JAVASCRIPT:
            formattedArgs = append(formattedArgs, fmt.Sprintf("%s: '%s'", name, value))
        case LUA:
            formattedArgs = append(formattedArgs, fmt.Sprintf("%s = '%s'", name, value))
        case RUBY:
            formattedArgs = append(formattedArgs, fmt.Sprintf("%s: '%s'", name, value))
        }
    }

    // Join the formatted arguments into a single string
    return strings.Join(formattedArgs, ", ")
}
