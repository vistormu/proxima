package evaluator

import (
    "proxima/ast"
    "proxima/builtins"
)

func Eval(node ast.Node) string {
    switch node := node.(type) {
    case *ast.Document:
        return evalDocument(node)
    case *ast.Paragraph:
        return evalParagraph(node)
    case *ast.Text:
        return evalText(node)
    case *ast.Tag:
        return evalTag(node)
    default:
        return ""
    }
}

func evalDocument(document *ast.Document) string {
    var result string

    for _, paragraph := range document.Paragraphs {
        result += Eval(paragraph) + "\n"
    }

    return result
}
func evalParagraph(paragraph *ast.Paragraph) string {
    var result string

    for _, inline := range paragraph.Content {
        result += Eval(inline)
    }

    return result
}

func evalText(text *ast.Text) string {
    return text.Content
}

func evalTag(tag *ast.Tag) string {
    var result string

    function, ok := builtins.Builtins[tag.Name]
    if !ok {
        return ""
    }
    for _, inline := range tag.Content {
        result += Eval(inline)
    }

    result = function(result)

    return result
}
