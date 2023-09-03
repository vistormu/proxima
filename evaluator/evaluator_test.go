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

    @center
    This is a centered paragraph.

    @right
    And this one is right-aligned.

    @center
    This one is centered again, but it has @bold{bold text} in it.
    `
    expected := `<div class="h1">
    This is a section title
</div>

<div class="paragraph">
    This is a paragraph.
</div>

<div class="paragraph center">
    This is a centered paragraph.
</div>

<div class="paragraph right">
    And this one is right-aligned.
</div>

<div class="paragraph center">
    This one is centered again, but it has <strong>bold text</strong> in it.
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
