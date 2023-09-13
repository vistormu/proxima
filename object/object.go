package object

type ObjectType int
const (
    STRING_OBJ ObjectType = iota
    ERROR_OBJ
)

type Object interface {
    Type() ObjectType
    Inspect() string
}

type String struct {
    Value string
}
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string { return s.Value }

type Error struct {
    Message string
}
func (e *Error) Type() ObjectType { return ERROR_OBJ }
func (e *Error) Inspect() string { return e.Message }
