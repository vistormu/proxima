package evaluator

import (
    "proxima/error"
    "proxima/ast"
    "proxima/builtins"
)

// TMP
const (
    preamble = `<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <style>
        @import url('https://fonts.googleapis.com/css2?family=Lora&family=Roboto&display=swap');
        .paragraph {
            font-family: "Lora", serif;
            font-size: 12pt;
            margin-top: 6px;
            margin-bottom: 6px;
            text-indent: 12px;
            text-align: justify;
            line-height: 1.25;
        }
        .h1 {
            margin-top: 32px;
            margin-bottom: 12px;
            font-size: 24px;
            font-weight: bold;
            font-family: "Roboto", sans-serif;
        }
        .h2 {
            margin-top: 24px;
            margin-bottom: 12px;
            font-size: 20px;
            font-weight: bold;
            font-family: "Roboto", sans-serif;
        }
        .h3 {
            margin-top: 20px;
            margin-bottom: 12px;
            font-size: 18px;
            font-weight: bold;
            font-family: "Roboto", sans-serif;
        }
        #center {
            text-align: center;
        }
        #right {
            text-align: right;
        }
        #monospace {
            font-family: monospace;
        }
    </style>
</head>
<body>
    `
    postamble = `</body>
</html>
`
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
    case *ast.Comment:
        return ""
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
        result += e.Eval(paragraph)
    }

    result = preamble + result + postamble

    return result
}
func (e *Evaluator) evalParagraph(paragraph *ast.Paragraph) string {
    var result string

    _, isText := paragraph.Content[0].(*ast.Text)
    _, isTag := paragraph.Content[0].(*ast.Tag)
    isBracketedTag := false
    if isTag && paragraph.Content[0].(*ast.Tag).Type == ast.BRACKETED {
        isBracketedTag = true
    }

    for _, inline := range paragraph.Content {
        result += e.Eval(inline)
    }

    if isText || isBracketedTag {
        result = "<div class=\"paragraph\">\n\t" + result + "\n</div>\n"
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
        e.addError("Unknown tag")
        return "??"
    }
    for _, inline := range tag.Arguments[0] {
        result += e.Eval(inline)
    }

    result = function(result, tag.Type)
    if result == "" {
        e.addError("Tag function returned empty string")
        return "??"
    }

    return result
}
