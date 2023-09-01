package evaluator

import (
	"proxima/parser"
	"testing"
)

func TestEvalText(t *testing.T) {
    tests := []struct { 
        input string
        expected string
    }{
        {"Hello, World!", "<p>Hello, World!</p>\n"},
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
        {"@bold{Hello, World!}", "<p><b>Hello, World!</b></p>\n"},
    }

    for _, test := range tests {
        evaluated := testEval(test.input)
        testString(t, evaluated, test.expected)
    }
}

func TestTagWrap(t *testing.T) {
    tests := []struct { 
        input string
        expected string
    }{
        {`@center
        This is centered text`,
        "<p><center>This is centered text</center></p>\n"},
    }

    for _, test := range tests {
        evaluated := testEval(test.input)
        testString(t, evaluated, test.expected)
    }
}

// HELPERS
func testEval(input string) string {
    parser := parser.New(input)
    document := parser.Parse()
    ev := New()

    return ev.Eval(document)
}
func testString(t *testing.T, input string, expected string) {
    evaluated := testEval(input)

    if evaluated != expected {
        t.Errorf("Expected %q, got %q", expected, evaluated)
    }
}
