package parser

import (
    "testing"
    "proxima/ast"
)

func TestParseParagraph(t *testing.T) {
    input := `@center 
    This is centered text!
    `
    p := New(input)
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should have 1 paragraph, got %d", len(document.Paragraphs))
    }
    
    tag, ok := document.Paragraphs[0].Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should have a tag, got %T", document.Paragraphs[0].Content[0])
    }
    if tag.Name != "center" {
        t.Fatalf("tag name should be 'center', got %s", tag.Name)
    }

    text, ok := document.Paragraphs[0].Content[1].(ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[0].Content[1])
    }
    if text.Content != "This is centered text!" {
        t.Fatalf("text should be 'This is centered text!', got %s", text.Content)
    }
}

func TestParseText(t *testing.T) {
    input := `This is text!`
    p := New(input)
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should have 1 paragraph, got %d", len(document.Paragraphs))
    }

    text, ok := document.Paragraphs[0].Content[0].(ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[0].Content[0])
    }
    if text.Content != "This is text!" {
        t.Fatalf("text should be 'This is text!', got %s", text.Content)
    }
}

func TestParseTagWithContent(t *testing.T) {
    input := `@bold{This is bold text!}`
    p := New(input)
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should have 1 paragraph, got %d", len(document.Paragraphs))
    }

    tag, ok := document.Paragraphs[0].Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should have a tag, got %T", document.Paragraphs[0].Content[0])
    }
    if tag.Name != "bold" {
        t.Fatalf("tag name should be 'bold', got %s", tag.Name)
    }

    text, ok := tag.Content[0].(ast.Text)
    if !ok {
        t.Fatalf("tag should have text, got %T", tag.Content[0])
    }
    if text.Content != "This is bold text!" {
        t.Fatalf("text should be 'This is bold text!', got %s", text.Content)
    }
}

func TestTagAndTagWithContent(t *testing.T) {
    input := `
    @center
    This is centered and @bold{bold} text!
    `

    p := New(input)
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 1 {
        t.Fatalf("document should have 1 paragraph, got %d", len(document.Paragraphs))
    }

    tag, ok := document.Paragraphs[0].Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should have a tag, got %T", document.Paragraphs[0].Content[0])
    }
    if tag.Name != "center" {
        t.Fatalf("tag name should be 'center', got %s", tag.Name)
    }

    text, ok := document.Paragraphs[0].Content[1].(ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[0].Content[1])
    }
    if text.Content != "This is centered and " {
        t.Fatalf("text should be 'This is centered and ', got %s", text.Content)
    }

    tag, ok = document.Paragraphs[0].Content[2].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should have a tag, got %T", document.Paragraphs[0].Content[2])
    }
    if tag.Name != "bold" {
        t.Fatalf("tag name should be 'bold', got %s", tag.Name)
    }

    text, ok = tag.Content[0].(ast.Text)
    if !ok {
        t.Fatalf("tag should have text, got %T", tag.Content[0])
    }
    if text.Content != "bold" {
        t.Fatalf("text should be 'bold', got %s", text.Content)
    }

    text, ok = document.Paragraphs[0].Content[3].(ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[0].Content[3])
    }
    if text.Content != " text!" {
        t.Fatalf("text should be ' text!', got %s", text.Content)
    }
}

func TestMultiParagraph(t *testing.T) {
    input := `
    This is the first paragraph.

    This is the second paragraph.
    `
    p := New(input)
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 2 {
        t.Fatalf("document should have 2 paragraphs, got %d", len(document.Paragraphs))
    }

    text, ok := document.Paragraphs[0].Content[0].(ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[0].Content[0])
    }
    if text.Content != "This is the first paragraph." {
        t.Fatalf("text should be 'This is the first paragraph.', got %s", text.Content)
    }

    text, ok = document.Paragraphs[1].Content[0].(ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[1].Content[0])
    }
    if text.Content != "This is the second paragraph." {
        t.Fatalf("text should be 'This is the second paragraph.', got %s", text.Content)
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
