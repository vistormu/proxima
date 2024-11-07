package evaluator

import (
	"fmt"
    "strings"
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

    textReplacements map[string]string

    interp Interpreter
}

// PUBLIC
func New(expressions []ast.Expression, file string, config *config.Config) (*Evaluator, error) {
    // load components
    components, err := loadComponents(expressions, config)
    if err != nil {
        return nil, err
    }

    textReplacements := make(map[string]string, len(config.Evaluator.TextReplacements))
    for _, replacement := range config.Evaluator.TextReplacements {
        textReplacements[replacement.From] = replacement.To
    }

    return &Evaluator{
        expressions: expressions,
        components: components,
        file: file,
        textReplacements: textReplacements,
        interp: NewInterpreter(config),
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
        return e.evaluateText(expression)
    case *ast.Tag:
        return e.evaluateTag(expression)
    default:
        return "", nil
    }
}

func (e *Evaluator) evaluateText(text *ast.Text) (string, error) {
    textValue := text.Value
    for from, to := range e.textReplacements {
        if !strings.Contains(textValue, from) {
            continue
        }
        textValue = strings.ReplaceAll(textValue, from, to)
    }

    return textValue, nil
}

func (e *Evaluator) evaluateTag(tag *ast.Tag) (string, error) {
    component := e.components[tag.Name]

    args := make([]struct{
        Name string
        Value string
    }, 0, len(tag.Args))

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
        
        args = append(args, struct{
            Name string
            Value string
        }{
            Name: name,
            Value: evaluatedArg,
        })
    }

    sort.Slice(args, func(i, j int) bool {
        return args[i].Name < args[j].Name
    })

    // interpret
    output, err := e.interp.Evaluate(args, component)
    if err != nil {
        return "", errors.NewEvalError(errors.ERROR_EXECUTING_SCRIPT, e.file, e.currentLine, component.name, err)
    }

    return output, nil
}
