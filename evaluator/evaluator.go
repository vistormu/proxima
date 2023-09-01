package evaluator

import (
    "proxima/error"
    "proxima/ast"
    "proxima/builtins"
)

type Evaluator struct {
    Errors []error.Error
}


// PUBLIC
func New() *Evaluator {
    return &Evaluator{}
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
        e.addError("Unknown node type")
        return ""
    }
}

// ERRORS
func (e *Evaluator) addError(msg string) {
    e.Errors = append(e.Errors, error.Error{
        Stage: "evaluator",
        Message: msg,
    })
}

// EVALUATION
func (e *Evaluator) evalDocument(document *ast.Document) string {
    var result string

    for _, paragraph := range document.Paragraphs {
        result += `<p>` + e.Eval(paragraph) + `</p>` + "\n"
    }

    return result
}
func (e *Evaluator) evalParagraph(paragraph *ast.Paragraph) string {
    var result string

    for _, inline := range paragraph.Content {
        result += e.Eval(inline)
    }

    return result
}

func (e *Evaluator) evalText(text *ast.Text) string {
    return text.Content
}

func (e *Evaluator) evalTag(tag *ast.Tag) string {
    var result string

    function, ok := builtins.Builtins[tag.Name]
    if !ok {
        return ""
    }
    for _, inline := range tag.Content {
        result += e.Eval(inline)
    }

    result = function(result)

    return result
}
