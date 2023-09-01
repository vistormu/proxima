package evaluator

import (
    "testing"
    "proxima/parser"
)

func testEval(input string) string {
    parser := parser.New(input)
    document := parser.Parse()

    return Eval(document)
}

func testString(t *testing.T, input string, expected string) {
    evaluated := testEval(input)

    if evaluated != expected {
        t.Errorf("Expected %q, got %q", expected, evaluated)
    }
}

func TestEvalText(t *testing.T) {
    tests := []struct { 
        input string
        expected string
    }{
        {"Hello, World!", "Hello, World!"},
    }

    for _, test := range tests {
        testString(t, test.input, test.expected)
    }
}

func TestEvalTag(t *testing.T) {
    tests := []struct { 
        input string
        expected string
    }{
        {`@h1{Hello, World!}`, `<h1>Hello, World!</h1>`},
    }

    for _, test := range tests {
        evaluated := testEval(test.input)
        testString(t, evaluated, test.expected)
    }
}
