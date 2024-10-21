package errors

type ErrorType int
const (
    // Cli errors
    WRONG_N_ARGS ErrorType = iota
    UNKNOWN_COMMAND
    INVALID_FILE_EXTENSION
    UNKNOWN_FLAG
    NO_OUTPUT_FLAG
    MISSING_FLAG_VALUE
    FILE_CREATION_ERROR

    // Parse errors
    EXPECTED_TOKEN
    UNEXPECTED_TOKEN
    WRONG_TAG_NAME
    UNCLOSED_TAG
    
    // Eval errors
    ERROR_EXECUTING_SCRIPT

    // Component errors
    MISSING_COMPONENTS
    ERROR_READING_DIR
    ERROR_READING_FILE
    DUPLICATE_COMPONENT
)
