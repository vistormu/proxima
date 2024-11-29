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
    // CLI errors
    N_ARGS      CliError = "wrong number of arguments\x1b[0m\n   |> expected: %v \n   |> got: %v"
    COMMAND     CliError = "unknown command\x1b[0m\n   |> got: %v\n   |> did you mean \"\x1b[35m%v\x1b[0m\"?"
    EXTENSION   CliError = "invalid file extension\x1b[0m\n   |> got: %v\n   |> expected: .prox"
    FLAG        CliError = "unknown flag\x1b[0m\n   |> got: %v\n   |> did you mean \"\x1b[35m%v\x1b[0m\"?"
    FLAG_VALUE  CliError = "no flag value provided\x1b[0m\n   |> flag: %v"
    OUTPUT_FLAG CliError = "no output flag provided\x1b[0m"
    CREATE_FILE CliError = "error creating file\x1b[0m\n   |> file: %v"
    READ_FILE   CliError = "error reading file\x1b[0m\n   |> file: %v"

    // Config errors
    CONFIG      ConfigError = "error reading config file\x1b[0m\n   |> displaying error:\n\n%v"
    
    // Parser errors
    EXPECTED_TOKEN   ParseError = "expected token\x1b[0m\n   |> file: %v\n   |> line: %v\n   |> expected: %v\n   |> got: %v"
    WRONG_TAG_NAME   ParseError = "wrong tag name\x1b[0m\n   |> file: %v\n   |> line: %v\n   |> tag: %v\n   |> tag names cannot contain spaces"
    UNCLOSED_TAG     ParseError = "unclosed tag\x1b[0m\n   |> file: %v\n   |> line: %v"

    // Evaluator errors
    SCRIPT EvalError = "error executing python component\x1b[0m\n   |> file: %v\n   |> line: %v\n   |> component: %v\n   |> full error:\n\n%v"

    // Component errors
    MISSING    ComponentError = "missing component definitions\x1b[0m\n   |> components: %v"
    READ_DIR   ComponentError = "error reading directory\x1b[0m\n   |> directory: %v\n   |> full error:\n\n%v"
    READ_FILE2 ComponentError = "error reading file\x1b[0m\n   |> file: %v\n   |> full error:\n\n%v"
    DUPLICATE  ComponentError = "duplicate component\x1b[0m\n   |> component: %v\n   |> to fix this, try using the \"\x1b[35muse_modules\x1b[0m\" option"
)


func (e CliError) String() string { return string(e) }
func (e ConfigError) String() string { return string(e) }
func (e ParseError) String() string { return string(e) }
func (e EvalError) String() string { return string(e) }
func (e ComponentError) String() string { return string(e) }

var stageMessages = map[reflect.Type]string{
    reflect.TypeOf(CliError("")): "|CLI error| ",
    reflect.TypeOf(ConfigError("")): "|Config error| ",
    reflect.TypeOf(ParseError("")): "|Parser error| ",
    reflect.TypeOf(EvalError("")): "|Evaluator error| ",
    reflect.TypeOf(ComponentError("")): "|Component error| ",
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
