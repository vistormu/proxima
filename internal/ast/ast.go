package ast

type Expression interface {
    Line() int
}

type Tag struct {
    Name string
    Args []Argument
    LineNumber int
}
func (t *Tag) Line() int {
    return t.LineNumber
}

type Argument struct {
    Name string
    Values []Expression
}

type Text struct {
    Value string
    LineNumber int
}
func (t *Text) Line() int {
    return t.LineNumber
}
