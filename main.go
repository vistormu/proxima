package main

import (
    "os"
    "proxima/commands"
)

func main() {
    commands.Execute(os.Args)
}
