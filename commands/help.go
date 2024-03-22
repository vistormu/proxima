package commands

import (
    "fmt"
)

func help(args []string) {
    if len(args) == 0 {
        helpBasic()
        return
    }
    switch args[0] {
    case "generate":
        helpGenerate()
    case "version":
        helpVersion()
    default:
        helpBasic()
    }
}

func helpBasic() {
    msg := `Usage:

    proxima <command> [arguments]

Commands:
    
    generate   Generate a new HTML file from a .prox file
    version    Display the current version

Use "proxima help <command>" for more information about a command.
`
    fmt.Println(msg)
}

func helpGenerate() {
    msg := `Usage:

    proxima generate [-c <components_path>] [arguments]

Description:

    Generate a new HTML file from a .prox file

Flags:

    -c <components_path>    Path to the components directory. By default, the components directory is "./components"

Arguments:

    filename.prox   The .prox file to generate an HTML file from
    all <dir>       Generate HTML files from all .prox files in the specified directory
`

    fmt.Println(msg)
}

func helpVersion() {
    msg := `Usage:

    proxima version

Description:

    Display the current version
`
    
    fmt.Println(msg)
}
