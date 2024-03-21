package evaluator

import (
	"fmt"
	"proxima/ast"
	"proxima/components"
	"proxima/error"
	"proxima/object"
	"strings"
)

type Evaluator struct {
    Errors []error.Error
    File string
}

// PUBLIC
func New(file string) *Evaluator {
    return &Evaluator{File: file}
}
func (e *Evaluator) Eval(node ast.Node) string {
    switch node := node.(type) {
    case *ast.Document:
        return e.evalDocument(node)
    case *ast.Paragraph:
        return e.evalParagraph(node)
    case *ast.Text:
        return e.evalText(node)
    case *ast.Tag:
        return e.evalTag(node)
    default:
        e.addError(fmt.Sprintf("unknown node type: %T", node), node.Line())
        return ""
    }
}

// ERRORS
func (e *Evaluator) addError(msg string, line int) {
    e.Errors = append(e.Errors, error.Error{
        Stage: "evaluator",
        Message: msg,
        Line: line,
        File: e.File,
    })
}

// EVALUATION
func (e *Evaluator) evalDocument(document *ast.Document) string {
    var result string

    for _, paragraph := range document.Paragraphs {
        result += e.Eval(paragraph)
    }

    return result
}
func (e *Evaluator) evalParagraph(paragraph *ast.Paragraph) string {
    var result string

    _, isText := paragraph.Content[0].(*ast.Text)
    _, isTag := paragraph.Content[0].(*ast.Tag)
    isBracketedTag := false
    if isTag {
        isBracketedTag = true
    }

    for _, inline := range paragraph.Content {
        result += e.Eval(inline)
    }

    if (isText || isBracketedTag) && !strings.HasPrefix(result, "<") {
        result = "<p>\n\t" + result + "\n</p>"
    }

    return result + "\n"
}

func (e *Evaluator) evalText(text *ast.Text) string {
    return text.Content
}

func (e *Evaluator) evalTag(tag *ast.Tag) string {
    function := getTagFuntion(tag.Name)
    if function == nil {
        e.addError("unknown tag: " + tag.Name, tag.Line())
        return ""
    }

    var evaluatedArguments []string
    for _, argument := range tag.Arguments {
        var evaluatedArgument string
        for _, inline := range argument {
            evaluatedArgument += e.Eval(inline)
        }
        evaluatedArguments = append(evaluatedArguments, evaluatedArgument)
    }

    result := function(evaluatedArguments)
    if result.Type() == object.ERROR_OBJ {
        e.addError(result.Inspect(), tag.Line())
        return ""
    }

    return result.Inspect()
}

func getTagFuntion(key string) components.ComponentFunction {
    if components.Components != nil {
        if function, ok := components.Components[key]; ok {
            return function
        }
    }

    return nil
}
