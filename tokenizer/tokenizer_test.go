package tokenizer

import (
    "testing"
    "visml/token"
)


func test(t *testing.T, input string, tests []struct { expectedType token.TokenType; expectedSubType token.TokenType; expectedContent string }) {
    tokenizer := New(input)

    for i, test := range tests {
        tok := tokenizer.GetToken()

        if tok.Type != test.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong. expected=%s, got=%s",
                i, token.TypeToString[test.expectedType], token.TypeToString[tok.Type])
        }

        if tok.Subtype != test.expectedSubType {
            t.Fatalf("tests[%d] - tokensubtype wrong. expected=%s, got=%s",
                i, token.TypeToString[test.expectedSubType], token.TypeToString[tok.Subtype])
        }

        if tok.Content != test.expectedContent {
            t.Fatalf("tests[%d] - content wrong. expected=%s, got=%s",
                i, test.expectedContent, tok.Content)
        }
    }
}

func TestTag(t *testing.T) {
    input := `@center This is centered text!`

    tests := []struct {
        expectedType token.TokenType
        expectedSubType token.TokenType
        expectedContent string
    }{
        {token.TAG, token.CENTER, ""},
        {token.TEXT, token.TEXT, "This is centered text!"},
        {token.CHAR, token.EOF, ""},
    }

    test(t, input, tests)
}

func TestTagWithContent(t *testing.T) {
    input := `@bold{This is bold text!}`

    tests := []struct {
        expectedType token.TokenType
        expectedSubType token.TokenType
        expectedContent string
    }{
        {token.TAG, token.BOLD, ""},
        {token.CHAR, token.LBRACE, ""},
        {token.TEXT, token.TEXT, "This is bold text!"},
        {token.CHAR, token.RBRACE, ""},
        {token.CHAR, token.EOF, ""},
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
        expectedSubType token.TokenType
        expectedContent string
    }{
        {token.CHAR, token.LINEBREAK, ""},
        {token.TAG, token.CENTER, ""},
        {token.CHAR, token.LINEBREAK, ""},
        {token.TEXT, token.TEXT, "This is centered and "},
        {token.TAG, token.BOLD, ""},
        {token.CHAR, token.LBRACE, ""},
        {token.TEXT, token.TEXT, "bold"},
        {token.CHAR, token.RBRACE, ""},
        {token.TEXT, token.TEXT, " text!"},
    }

    test(t, input, tests)
}
