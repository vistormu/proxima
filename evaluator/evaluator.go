package evaluator

import (
    "proxima/ast"
)

func Eval(node ast.Node) string {
    switch node := node.(type) {
    case *ast.Document:
        return evalDocument(node)

    default:
        return ""
    }
}

func evalDocument(document *ast.Document) string {
    var result string

    for _, paragraph := range document.Paragraphs {
        result = Eval(paragraph)
    }

    return result
}
