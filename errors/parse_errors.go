package errors

import (
    "fmt"
)

type ParseError struct {
    file string
    line int
    message string
}

func NewParseError(errorType ErrorType, file string, line int, args ...any) error {
    var message string

    switch errorType {
    case EXPECTED_TOKEN:
        message = fmt.Sprintf("expected %v, received %v", args[0], args[1])

    case UNEXPECTED_TOKEN:
        message = fmt.Sprintf("unexpected token: %v", args[0])

    case WRONG_TAG_NAME:
        message = fmt.Sprintf("wrong tag name: %v", args[0])

    case UNCLOSED_TAG:
        message = "unclosed tag"
    }

    return ParseError{file, line, message}
}

func (e ParseError) Error() string {
    return fmt.Sprintf("\x1b[31mproxima parser error in file \"%s\", line %d:\x1b[0m\n-> %s\n", e.file, e.line, e.message)
}
