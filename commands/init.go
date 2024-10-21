package commands


import (
    "os"
    "proxima/config"
    "proxima/errors"
)


func init_(args []string) error {
    if len(args) != 0 {
        return errors.NewCliError(errors.WRONG_N_ARGS, 0, len(args))
    }

    // create proxima.toml
    f, err := os.Create("proxima.toml")
    if err != nil {
        return errors.NewCliError(errors.FILE_CREATION_ERROR, "proxima.toml")
    }
    defer f.Close()

    _, err = f.WriteString(config.DefaultConfig)
    if err != nil {
        return errors.NewCliError(errors.FILE_CREATION_ERROR, "proxima.toml")
    }

    // create components directory
    err = os.Mkdir("components", 0755)
    if err != nil {
        return errors.NewCliError(errors.FILE_CREATION_ERROR, "components")
    }

    // create a python component
    f, err = os.Create("components/proxima.py")
    if err != nil {
        return errors.NewCliError(errors.FILE_CREATION_ERROR, "components/proxima.py")
    }
    defer f.Close()

    _, err = f.WriteString("def proxima() -> str:\n    return \"hello from proxima!\"\n")

    // create main file
    f, err = os.Create("main" + MAIN_EXT)
    if err != nil {
        return errors.NewCliError(errors.FILE_CREATION_ERROR, "main" + MAIN_EXT)
    }
    defer f.Close()

    _, err = f.WriteString("@proxima{}")
    if err != nil {
        return errors.NewCliError(errors.FILE_CREATION_ERROR, "main" + MAIN_EXT)
    }

    return nil
}
