package commands

import (
    "fmt"
    "strings"
)

const (
    MAIN_EXT = ".prox"
    VERSION = "0.2.0"
)

func Execute(args []string) {
    switch args[0] {
    case "version":
        version()
    case "generate":
        args = args[1:]

        // flags
        var inputFiles []string
        var componentsPath string

        for i, arg := range args {
            if arg == "-c" && i + 1 < len(args) {
                componentsPath = args[i + 1]
                args = append(args[:i], args[i + 2:]...)
                break
            }
        }
        inputFiles = args
        
        // input files
        if len(inputFiles) == 0 {
            exitOnError("usage: proxima <filename>.prox or proxima all to process all .prox files")
        }
        if inputFiles[0] == "all" {
            inputFiles = getAllFiles("./")
        }

        // components flag
        if componentsPath != "" && !dirExists(componentsPath) {
            exitOnError("components directory does not exist")
        }

        // generate html files
        for _, file := range inputFiles {
            if !strings.HasSuffix(file, MAIN_EXT) {
                exitOnError(fmt.Sprintf("file %s is not a .prox file", file))
            }
            generate(file, componentsPath)
        }
    case "watch":
        watch()
    default:
        help(args[1:])
    }
}
