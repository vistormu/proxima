package parser

import (
    "testing"
    "proxima/ast"
)

func TestParser(t *testing.T) {
    input := `
    # This is a comment

    This is the first paragraph.

    @center
    This is the second paragraph.

    This is the third paragraph with @bold{bold text}.
    `
    p := New(input)
    document := p.Parse()

    checkParserErrors(t, p)

    if len(document.Paragraphs) != 3 {
        t.Fatalf("document should have 3 paragraphs, got %d", len(document.Paragraphs))
    }

    text, ok := document.Paragraphs[0].Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[0].Content[0])
    }
    if text.Content != "This is the first paragraph." {
        t.Fatalf("text should be 'This is the first paragraph.', got %s", text.Content)
    }

    tag, ok := document.Paragraphs[1].Content[0].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should have a tag, got %T", document.Paragraphs[1].Content[0])
    }
    if tag.Name != "center" {
        t.Fatalf("tag name should be 'center', got %s", tag.Name)
    }

    text, ok = tag.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should have text, got %T", tag.Content[0])
    }
    if text.Content != "This is the second paragraph." {
        t.Fatalf("text should be 'This is the second paragraph.', got %s", text.Content)
    }

    text, ok = document.Paragraphs[2].Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[2].Content[0])
    }
    if text.Content != "This is the third paragraph with " {
        t.Fatalf("text should be 'This is the third paragraph with ', got %s", text.Content)
    }

    tag, ok = document.Paragraphs[2].Content[1].(*ast.Tag)
    if !ok {
        t.Fatalf("paragraph should have a tag, got %T", document.Paragraphs[2].Content[1])
    }
    if tag.Name != "bold" {
        t.Fatalf("tag name should be 'bold', got %s", tag.Name)
    }
    
    text, ok = tag.Content[0].(*ast.Text)
    if !ok {
        t.Fatalf("tag should have text, got %T", tag.Content[0])
    }
    if text.Content != "bold text" {
        t.Fatalf("text should be 'bold text', got %s", text.Content)
    }

    text, ok = document.Paragraphs[2].Content[2].(*ast.Text)
    if !ok {
        t.Fatalf("paragraph should have text, got %T", document.Paragraphs[2].Content[2])
    }
    if text.Content != "." {
        t.Fatalf("text should be '.', got %s", text.Content)
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
