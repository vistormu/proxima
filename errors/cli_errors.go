package errors

import (
    "fmt"
)

type CliError struct {
    message string
}

func NewCliError(errorType ErrorType, args ...any) error {
    var message string

    switch errorType {
    case WRONG_N_ARGS:
        message = fmt.Sprintf("expected %v arguments, received %v", args[0], args[1])

    case UNKNOWN_COMMAND:
        message = fmt.Sprintf("unknown command: %s", args[0].(string))

    case INVALID_FILE_EXTENSION:
        message = fmt.Sprintf("invalid file extension %s, expected .prox file", args[0].(string))

    case UNKNOWN_FLAG:
        message = fmt.Sprintf("unknown flag: %s", args[0].(string))

    case NO_OUTPUT_FLAG:
        message = "no output flag provided"

    case MISSING_FLAG_VALUE:
        message = fmt.Sprintf("missing value for flag: %s", args[0].(string))

    case FILE_CREATION_ERROR:
        message = fmt.Sprintf("error creating file \"%s\"", args[0].(string))
    }

    return CliError{message}
}

func (e CliError) Error() string {
    return fmt.Sprintf("\x1b[31mproxima CLI error\x1b[0m\n-> %s\n", e.message)
}
