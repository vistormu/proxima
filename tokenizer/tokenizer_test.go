package tokenizer

import (
    "testing"
    "proxima/token"
)


func test(t *testing.T, input string, tests []struct { expectedType token.TokenType; expectedContent string }) {
    tokenizer := New(input)

    for i, test := range tests {
        tok := tokenizer.GetToken()

        if tok.Type != test.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
                i, token.TypeToString[test.expectedType], token.TypeToString[tok.Type])
        }

        if tok.Literal != test.expectedContent {
            t.Fatalf("tests[%d] - content wrong. expected=%s, got=%s",
                i, test.expectedContent, tok.Literal)
        }
    }
}

func TestTag(t *testing.T) {
    input := `@center`

    tests := []struct {
        expectedType token.TokenType
        expectedContent string
    }{
        {token.TAG, "@center"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}

func TestTagWithContent(t *testing.T) {
    input := `@bold{This is bold text!}`

    tests := []struct {
        expectedType token.TokenType
        expectedContent string
    }{
        {token.TAG, "@bold"},
        {token.LBRACE, "{"},
        {token.TEXT, "This is bold text!"},
        {token.RBRACE, "}"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}

func TestTagsAndTagsWithContent(t *testing.T) {
    input := `
@center
This is centered and @bold{bold} text!
`

    tests := []struct {
        expectedType token.TokenType
        expectedContent string
    }{
        {token.LINEBREAK, "\n"},
        {token.TAG, "@center"},
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is centered and "},
        {token.TAG, "@bold"},
        {token.LBRACE, "{"},
        {token.TEXT, "bold"},
        {token.RBRACE, "}"},
        {token.TEXT, " text!"},
        {token.LINEBREAK, "\n"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}

func TestMultiParagraph(t *testing.T) {
    input := `
This is the first paragraph.

@center
This is the second paragraph.

This is the third paragraph with @bold{bold text}.
`
    tests := []struct {
        expectedType token.TokenType
        expectedContent string
    }{
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is the first paragraph."},
        {token.LINEBREAK, "\n"},
        {token.LINEBREAK, "\n"},
        {token.TAG, "@center"},
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is the second paragraph."},
        {token.LINEBREAK, "\n"},
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is the third paragraph with "},
        {token.TAG, "@bold"},
        {token.LBRACE, "{"},
        {token.TEXT, "bold text"},
        {token.RBRACE, "}"},
        {token.TEXT, "."},
        {token.LINEBREAK, "\n"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}

func TestComment(t *testing.T) {
    input := `
This is the first paragraph.

# This is a comment

This is the second paragraph.

# This is a
# Double line comment

This is the third paragraph.
`
    tests := []struct {
        expectedType token.TokenType
        expectedContent string
    }{
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is the first paragraph."},
        {token.LINEBREAK, "\n"},
        {token.LINEBREAK, "\n"},
        {token.HASH, "#"},
        {token.TEXT, " This is a comment"},
        {token.LINEBREAK, "\n"},
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is the second paragraph."},
        {token.LINEBREAK, "\n"},
        {token.LINEBREAK, "\n"},
        {token.HASH, "#"},
        {token.TEXT, " This is a"},
        {token.LINEBREAK, "\n"},
        {token.HASH, "#"},
        {token.TEXT, " Double line comment"},
        {token.LINEBREAK, "\n"},
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is the third paragraph."},
        {token.LINEBREAK, "\n"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}

func TestEscacpeCharacter(t *testing.T) {
    input := `text \@ \{ \}`

    tests := []struct {
        expectedType token.TokenType
        expectedContent string
    }{
        {token.TEXT, "text "},
        {token.TEXT, "@"},
        {token.TEXT, " "},
        {token.TEXT, "{"},
        {token.TEXT, " "},
        {token.TEXT, "}"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}
