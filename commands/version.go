package commands

import (
    "fmt"
    "proxima/errors"
)

func version(args []string) error {
    if len(args) > 0 {
        return errors.NewCliError(errors.WRONG_N_ARGS, 0, len(args))
    }

    msg := fmt.Sprintf("Proxima version: %s", VERSION)
    fmt.Println(msg)

    return nil
}
