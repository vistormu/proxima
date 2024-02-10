package error

import (
    "fmt"
)

type Error struct {
    File string
    Stage string
    Line int
    Message string
}
func (e *Error) String() string {
    return fmt.Sprintf("\x1b[31m -> |%s| %s (line %d)\x1b[0m \n %s", e.Stage, e.File, e.Line, e.Message)
}
