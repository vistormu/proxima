package errors

import (
    "fmt"
)

type EvalError struct {
    file string
    line int
    message string
}

func NewEvalError(errorType ErrorType, file string, line int, args ...any) error {
    var message string

    switch errorType {
    case ERROR_EXECUTING_SCRIPT:
        message = fmt.Sprintf("error executing function \"%v\"\n%v", args[0], args[1])
    }

    return EvalError{file, line, message}
}

func (e EvalError) Error() string {
    return fmt.Sprintf("\x1b[31mproxima evaluation error in file \"%s\", line %d:\x1b[0m\n-> %s\n", e.file, e.line, e.message)
}
