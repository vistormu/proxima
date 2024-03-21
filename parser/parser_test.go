package parser

import (
    "testing"
    "proxima/ast"
)

func TestText(t *testing.T) {
    input := `This is a paragraph of text.`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    text, ok := paragraph.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != "This is a paragraph of text." {
        t.Fatalf("text inline should contain 'This is a paragraph of text.', got '%s'", text.Content)
    }
}

func TestBracketedTag(t *testing.T) {
    input := `@bold{This is a paragraph of text.}`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    tag, ok := paragraph.Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "bold" {
        t.Fatalf("tag should be named 'bold', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[0]))
    }

    text, ok := tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "This is a paragraph of text." {
        t.Fatalf("text inline should contain 'This is a paragraph of text.', got '%s'", text.Content)
    }
}

func TestSelfClosingTag(t *testing.T) {
    input := `
@line

this is more text @rightarrow and more text
`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 2 {
        t.Fatalf("document should contain 2 paragraphs, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    tag, ok := paragraph.Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }
    
    if tag.Name != "line" {
        t.Fatalf("tag should be named 'line', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 0 {
        t.Fatalf("tag should contain 0 inlines, got %d", len(tag.Arguments))
    }

    paragraph = document.Paragraphs[1]
    if len(paragraph.Content) != 3 {
        t.Fatalf("paragraph should contain 3 inlines, got %d", len(paragraph.Content))
    }

    text, ok := paragraph.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != "this is more text " {
        t.Fatalf("text inline should contain 'this is more text ', got '%s'", text.Content)
    }

    tag, ok = paragraph.Content[1].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "rightarrow" {
        t.Fatalf("tag should be named 'rightarrow', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 0 {
        t.Fatalf("tag should contain 0 inlines, got %d", len(tag.Arguments))
    }

    text, ok = paragraph.Content[2].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != " and more text" {
        t.Fatalf("text inline should contain ' and more text', got '%s'", text.Content)
    }
}

func TestParseEscapeCharacter(t *testing.T) {
    input := `\@escape \@character`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 4 {
        t.Fatalf("paragraph should contain 4 inline, got %d", len(paragraph.Content))
    }

    text, ok := paragraph.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != "@" {
        t.Fatalf("text inline should contain '@', got '%s'", text.Content)
    }

    text_2, ok := paragraph.Content[1].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text_2.Content != "escape " {
        t.Fatalf("text inline should contain 'escape ', got '%s'", text_2.Content)
    }

    text_3, ok := paragraph.Content[2].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text_3.Content != "@" {
        t.Fatalf("text inline should contain '@', got '%s'", text_3.Content)
    }

    text_4, ok := paragraph.Content[3].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text_4.Content != "character" {
        t.Fatalf("text inline should contain 'character', got '%s'", text_4.Content)
    }
}

func TestTextAndTags(t *testing.T) {
    input := `This is a paragraph of text with @bold{bold text} and @italic{italic text}.`

    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 5 {
        t.Fatalf("paragraph should contain 5 inlines, got %d", len(paragraph.Content))
    }

    text, ok := paragraph.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != "This is a paragraph of text with " {
        t.Fatalf("text inline should contain 'This is a paragraph of text with ', got '%s'", text.Content)
    }

    tag, ok := paragraph.Content[1].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "bold" {
        t.Fatalf("tag should be named 'bold', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[0]))
    }

    text, ok = tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "bold text" {
        t.Fatalf("text inline should contain 'bold text', got '%s'", text.Content)
    }

    text, ok = paragraph.Content[2].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != " and " {
        t.Fatalf("text inline should contain ' and ', got '%s'", text.Content)
    }

    tag, ok = paragraph.Content[3].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "italic" {
        t.Fatalf("tag should be named 'italic', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 1 {
        t.Fatalf("tag should contain 1 argument, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[0]))
    }

    text, ok = tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "italic text" {
        t.Fatalf("text inline should contain 'italic text', got '%s'", text.Content)
    }

    text, ok = paragraph.Content[4].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != "." {
        t.Fatalf("text inline should contain '.', got '%s'", text.Content)
    }
}

func TestMultiArgumentTag( t *testing.T) {
    input := `@url{http://www.google.com}{Google}`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    tag, ok := paragraph.Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "url" {
        t.Fatalf("tag should be named 'url', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 2 {
        t.Fatalf("tag should contain 2 arguments, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[0]))
    }

    text, ok := tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "http://www.google.com" {
        t.Fatalf("text inline should contain 'http://www.google.com', got '%s'", text.Content)
    }

    if len(tag.Arguments[1]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[1]))
    }

    text, ok = tag.Arguments[1][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "Google" {
        t.Fatalf("text inline should contain 'Google', got '%s'", text.Content)
    }
}

func TestNestedBracketing(t *testing.T) {
    input := `@bold{bold and @italic{italic}}`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    tag, ok := paragraph.Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "bold" {
        t.Fatalf("tag should be named 'bold', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 1 {
        t.Fatalf("tag should contain 1 argument, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 2 {
        t.Fatalf("tag should contain 2 inlines, got %d", len(tag.Arguments[0]))
    }

    text, ok := tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "bold and " {
        t.Fatalf("text inline should contain 'bold and ', got '%s'", text.Content)
    }

    tag, ok = tag.Arguments[0][1].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "italic" {
        t.Fatalf("tag should be named 'italic', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 1 {
        t.Fatalf("tag should contain 1 argument, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[0]))
    }

    text, ok = tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "italic" {
        t.Fatalf("text inline should contain 'italic', got '%s'", text.Content)
    }
}

func TestMultiText(t *testing.T) {
    input := `
This is a paragraph of text.
This is the same paragraph of text.
`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    text, ok := paragraph.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should contain a text inline")
    }

    if text.Content != "This is a paragraph of text.\nThis is the same paragraph of text." {
        t.Fatalf("text inline should contain 'This is a paragraph of text.\nThis is the same paragraph of text.', got '%s'", text.Content)
    }
}

func TestEmptyArgument(t *testing.T) {
    input := `@bold{}`
    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    tag, ok := paragraph.Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "bold" {
        t.Fatalf("tag should be named 'bold', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 1 {
        t.Fatalf("tag should contain 1 argument, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 0 {
        t.Fatalf("tag should contain 0 inlines, got %d", len(tag.Arguments[0]))
    }
}

func TestMultipleClosedArguments(t *testing.T) {
    input := `@bold{first}{}{third}`

    p := New(input, "test")
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should contain 1 paragraph, got %d", len(document.Paragraphs))
    }

    paragraph := document.Paragraphs[0]
    if len(paragraph.Content) != 1 {
        t.Fatalf("paragraph should contain 1 inline, got %d", len(paragraph.Content))
    }

    tag, ok := paragraph.Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should contain a tag inline")
    }

    if tag.Name != "bold" {
        t.Fatalf("tag should be named 'bold', got '%s'", tag.Name)
    }

    if len(tag.Arguments) != 3 {
        t.Fatalf("tag should contain 3 arguments, got %d", len(tag.Arguments))
    }

    if len(tag.Arguments[0]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[0]))
    }

    text, ok := tag.Arguments[0][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "first" {
        t.Fatalf("text inline should contain 'first', got '%s'", text.Content)
    }

    if len(tag.Arguments[1]) != 0 {
        t.Fatalf("tag should contain 0 inlines, got %d", len(tag.Arguments[1]))
    }

    if len(tag.Arguments[2]) != 1 {
        t.Fatalf("tag should contain 1 inline, got %d", len(tag.Arguments[2]))
    }

    text, ok = tag.Arguments[2][0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should contain a text inline")
    }

    if text.Content != "third" {
        t.Fatalf("text inline should contain 'third', got '%s'", text.Content)
    }
}

func checkParserErrors(t *testing.T, p *Parser) {
    errors := p.Errors

    if len(errors) == 0 {
        return
    }

    t.Errorf("parser has %d errors", len(errors))
    for _, error := range errors {
        t.Errorf(error.String())
    }
    t.FailNow()
}
