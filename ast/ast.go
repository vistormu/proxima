package ast

import (
    "visml/token"
)

// Interfaces
type Node interface {}
type Inline interface {}

// Types
type Paragraph struct {
    Node
    Content []Inline
}

type Document struct {
    Paragraphs []Paragraph
}

type Text struct {
    Inline
    Content string
}

type Tag struct {
    Inline
    Type token.TokenType
    Content []Inline
}
