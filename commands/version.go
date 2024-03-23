package commands

import (
    "fmt"
)

func version() {
    msg := fmt.Sprintf("Proxima version: %s", VERSION)
    fmt.Println(msg)
}
