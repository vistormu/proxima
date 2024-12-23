package errors

import (
    "fmt"
    "strings"
    "reflect"
)


type ErrorType interface {
    String() string
}

type CliError string
type ConfigError string
type ParseError string
type EvalError string
type ComponentError string

const (
    E = "\x1b[0m"
    I = "\n   |> "
    F = "\n   |> full error:\n\n%v"
)


const (
    // CLI errors
    N_ARGS      CliError = "wrong number of arguments"+E+I+"expected: %v"+I+"got: %v"
    COMMAND     CliError = "unknown command"+E+I+"got: %v"+I+"did you mean \"\x1b[35m%v\x1b[0m\"?"
    EXTENSION   CliError = "invalid file extension"+E+I+"got: %v"+I+"expected: .prox"
    FLAG        CliError = "unknown flag"+E+I+"got: %v"+I+"did you mean \"\x1b[35m%v\x1b[0m\"?"
    FLAG_VALUE  CliError = "no flag value provided"+E+I+"flag: %v"
    OUTPUT_FLAG CliError = "no output flag provided"+E
    CREATE_FILE CliError = "error creating file"+E+I+"file: %v"
    READ_FILE   CliError = "error reading file"+E+I+"file: %v"

    // Config errors
    CONFIG ConfigError = "error reading config file"+E+F
    
    // Parser errors
    EXPECTED_TOKEN ParseError = "expected token"+E+I+"file: %v"+I+"line: %v"+I+"expected: %v"+I+"got: %v"
    WRONG_TAG_NAME ParseError = "wrong tag name"+E+I+"file: %v"+I+"line: %v"+I+"tag: %v"+I+"tag names cannot contain spaces"
    UNCLOSED_TAG   ParseError = "unclosed tag"+E+I+"file: %v"+I+"line: %v"

    // Evaluator errors
    SCRIPT           EvalError = "error executing python component"+E+I+"file: %v"+I+"line: %v"+I+"component: %v"+F
    NO_DEF           EvalError = "missing function definition"+E+I+"component: %v"
    NAME_MISMATCH    EvalError = "name mismatch"+E+I+"component: %v"+I+"file name: %v"+I+"function name: %v"+I+"function name and file name must match"
    INIT_INTERPRETER EvalError = "error initializing python interpreter"+E+F
    INTERPRETER_EVAL EvalError = "error evaluating python script"+E+F

    // Component errors
    MISSING    ComponentError = "missing component definitions"+E+I+"components: %v"
    READ_DIR   ComponentError = "error reading directory"+E+I+"directory: %v"+F
    READ_FILE2 ComponentError = "error reading file"+E+I+"file: %v"+F
    DUPLICATE  ComponentError = "duplicate component"+E+I+"component: %v"+I+"to fix this, try using the \"\x1b[35muse_modules\x1b[0m\" option"
)


func (e CliError) String() string { return string(e) }
func (e ConfigError) String() string { return string(e) }
func (e ParseError) String() string { return string(e) }
func (e EvalError) String() string { return string(e) }
func (e ComponentError) String() string { return string(e) }

var stageMessages = map[reflect.Type]string{
    reflect.TypeOf(CliError("")): "|cli error| ",
    reflect.TypeOf(ConfigError("")): "|config error| ",
    reflect.TypeOf(ParseError("")): "|parser error| ",
    reflect.TypeOf(EvalError("")): "|evaluator error| ",
    reflect.TypeOf(ComponentError("")): "|component error| ",
}

type Error struct {
    message string
}

func New(errorType ErrorType, args ...any) error {
    stageMessage := stageMessages[reflect.TypeOf(errorType)]
    errorMessage := errorType.String()

    message := "\x1b[31m-> " + stageMessage + errorMessage + "\n"
    n := strings.Count(message, "%v")

    if len(args) != n {
        panic(fmt.Sprintf("expected %v arguments, got %v", n, len(args)))
    }

    message = fmt.Sprintf(message, args...)

    return Error{message}
}

func (e Error) Error() string {
    return e.message
}
