package main

import (
    "os"
    "proxima/commands"
)

func main() {
    args := os.Args[1:]
    commands.Execute(args)
}
