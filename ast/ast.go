package ast

// Interfaces
type Node interface {
    Line() int
}
type Inline interface {
    Node
}

// Types
type Paragraph struct {
    Node
    Content []Inline
}
func (p *Paragraph) Line() int {
    return p.Content[0].Line()
}

type Document struct {
    Paragraphs []*Paragraph
}
func (d *Document) Line() int {
    return d.Paragraphs[0].Line()
}

type Text struct {
    Inline
    Content string
    LineNumber int
}
func (t *Text) Line() int {
    return t.LineNumber
}

type Tag struct {
    Inline
    Name string
    Arguments [][]Inline
    LineNumber int
}
func (t *Tag) Line() int {
    return t.LineNumber
}
