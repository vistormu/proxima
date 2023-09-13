package object

type ObjectType int
const (
    STRING ObjectType = iota
    ERROR
)

type Object interface {
    Type() ObjectType
    Inspect() string
}

type String struct {
    Value string
}
func (s *String) Type() ObjectType { return STRING }
func (s *String) Inspect() string { return s.Value }

type Error struct {
    Message string
}
func (e *Error) Type() ObjectType { return ERROR }
func (e *Error) Inspect() string { return e.Message }
