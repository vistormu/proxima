package evaluator

import (
	"proxima/parser"
	"testing"
)

func TestEval(t *testing.T) {
    input := `
    @h1
    This is a section title

    This is a paragraph.
    `
    expected := `<div class="h1">
    This is a section title
</div>
<div class="paragraph">
    This is a paragraph.
</div>
`
    testString(t, input, expected)
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
        t.Errorf("Expected:\n %q\n Got:\n %q", expected, evaluated)
    }
}
