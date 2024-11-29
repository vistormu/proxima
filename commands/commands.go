package commands

import (
    "fmt"
    "os"

    "proxima/errors"
)

const (
    MAIN_EXT = ".prox"
    VERSION = "0.5.0"
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
        os.Exit(1)
    }
}

func execute(args []string) error {
    if len(args) < 2 {
        return errors.New(errors.N_ARGS, "at least one", len(args)-1)
    }
    args = args[1:]

    command, ok := commands[args[0]]
    if !ok {
        closestMatch := findClosestMatch(args[0], keys(commands))
        return errors.New(errors.COMMAND, args[0], closestMatch)
    }

    return command(args[1:])
}
