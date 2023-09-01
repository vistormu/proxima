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
        {token.TAG, "center"},
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
        {token.TAG, "bold"},
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
        {token.TAG, "center"},
        {token.LINEBREAK, "\n"},
        {token.TEXT, "This is centered and "},
        {token.TAG, "bold"},
        {token.LBRACE, "{"},
        {token.TEXT, "bold"},
        {token.RBRACE, "}"},
        {token.TEXT, " text!"},
        {token.LINEBREAK, "\n"},
        {token.EOF, ""},
    }

    test(t, input, tests)
}