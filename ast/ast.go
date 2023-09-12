package ast

// Interfaces
type Node interface {}
type Inline interface {}

// Types
type Paragraph struct {
    Node
    Content []Inline
}

type Document struct {
    Paragraphs []*Paragraph
}

type Text struct {
    Inline
    Content string
}

type TagType int
const (
    WRAPPING TagType = iota
    BRACKETED
    SELF_CLOSING
)

type Tag struct {
    Inline
    Name string
    Type TagType
    Content []Inline
}

type Comment struct {
    Inline
    Content string
}
