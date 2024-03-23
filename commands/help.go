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
    case "watch":
        helpWatch()
    default:
        helpError()
    }
}

func helpError() {
    msg := `Error:

    Command not found

Use "proxima help" for more information
`
    fmt.Println(msg)
}

func helpBasic() {
    msg := `Usage:

    proxima <command> [arguments]

Commands:
    
    generate   Generate a new HTML file from a .prox file
    version    Display the current version
    watch      Watch a .prox file for changes and generate the corresponding HTML file

Use "proxima help <command>" for more information about a command.
`
    fmt.Println(msg)
}

func helpGenerate() {
    msg := `Usage:

    proxima generate [-c <components_path>] [-r] [arguments]

Description:

    Generate a new HTML file from a .prox file

Flags:

    -c <components_path>    Path to the components directory. By default, the components directory is "./components"

    -r                      Generate HTML files from all .prox files recursively

Arguments:
    
    file1 dir1 [file2 dir2] List of .prox files and directories to generate HTML files from

Examples:
    
    proxima generate -c ./components file1.prox file2.prox
    proxima generate -r dir1 dir2
    proxima generate dir1
    proxima generate file1.prox dir1
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

func helpWatch() {
    msg := `Usage:

    proxima watch [arguments]

Description:

    Watch a .prox file for changes and generate the corresponding HTML file

Arguments:
    
    file Watch a single .prox file

Examples:

    proxima watch file.prox
`

    fmt.Println(msg)
}
