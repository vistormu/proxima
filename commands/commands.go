package commands

import (
    "fmt"
    "proxima/errors"
)

const (
    MAIN_EXT = ".prox"
    VERSION = "0.4.0"
)

type CommnadFunc func(args []string) error
var commands = map[string]CommnadFunc{
    "init": init_,
    "make": make_,
    "help": help,
    "version": version,
}

func Execute(args []string) {
    err := execute(args)
    if err != nil {
        fmt.Println(err.Error())
    }
}

func execute(args []string) error {
    if len(args) < 2 {
        return errors.NewCliError(errors.WRONG_N_ARGS, "at least one", len(args)-1)
    }
    args = args[1:]

    command, ok := commands[args[0]]
    if !ok {
        return errors.NewCliError(errors.UNKNOWN_COMMAND, args[0])
    }

    return command(args[1:])
}
