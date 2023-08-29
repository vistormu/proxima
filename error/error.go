package error

import (
    "fmt"
)

type Error struct {
    Stage string
    Line int
    Message string
}
func (e *Error) String() string {
    return fmt.Sprintf("%s in line %d. %s", e.Stage, e.Line, e.Message)
}
