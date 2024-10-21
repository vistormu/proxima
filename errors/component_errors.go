package errors

import (
    "fmt"
)

type ComponentError struct {
    message string
}

func NewComponentError(errorType ErrorType, args ...any) error {
    var message string

    switch errorType {
    case MISSING_COMPONENTS:
        message = fmt.Sprintf("the following tags have a missing definition: %v", args[0])

    case ERROR_READING_DIR:
        message = fmt.Sprintf("error reading components directory \"%v\"\n%v", args[0], args[1])

    case ERROR_READING_FILE:
        message = fmt.Sprintf("error reading file \"%v\"\n%v", args[0], args[1])

    case DUPLICATE_COMPONENT:
        message = fmt.Sprintf("detected duplicate component \"%v\". choose a different name or use modules to give tags a namespace", args[0])
    }

    return ComponentError{message}
}

func (e ComponentError) Error() string {
    return fmt.Sprintf("\x1b[31mproxima component error:\x1b[0m\n-> %s\n", e.message)
}
