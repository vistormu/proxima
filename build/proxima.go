package main

import (
    "fmt"
    "io"
    "os"
    "strings"
    "path/filepath"
    "proxima/parser"
    "proxima/evaluator"
    "proxima/components"
)

const (
    MAIN_EXT = ".prox"
    PRE_HEAD = `<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta charset="UTF-8">
`
    POST_HEAD = `</head>
<body>
`
    POST_BODY = `</body>
</html>
`
)

func exitOnError(msg string) {
    fmt.Println("\x1b[31m -> |build| " + msg + "\x1b[0m")
    os.Exit(1)
}

func dirExists(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return info.IsDir()
}

func getAllFiles(dirPath string) []string {
    files := []string{}
    err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(info.Name(), MAIN_EXT) {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        exitOnError(err.Error())
    }
    return files
}

func main() {
    args := os.Args[1:]

    // flags
    var inputFiles []string
    var outputFilename string
    var componentsPath string

    for i, arg := range args {
        if arg == "-o" && i + 1 < len(args) {
            outputFilename = args[i + 1]
            args = append(args[:i], args[i + 2:]...)
            break
        }
    }

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

    // output flag
    if outputFilename != "" && !strings.HasSuffix(outputFilename, ".html") {
        exitOnError("output file must have .html extension")
    }
    if outputFilename != "" && len(inputFiles) > 1 {
        exitOnError("output file flag can only be used with one input file")
    }

    // components flag
    if componentsPath != "" && !dirExists(componentsPath) {
        exitOnError("components directory does not exist")
    }

    // generate html files
    for _, file := range inputFiles {
        generate(file, outputFilename, componentsPath)
    }
}

func generate(filename string, outputFilename string, componentsPath string) {
    if !strings.HasSuffix(filename, MAIN_EXT) {
        exitOnError(fmt.Sprintf("file %s is not a .prox file", filename))
    }
    splitFilename := strings.SplitN(filename, ".", 2)
    name := splitFilename[0]
    extension := "." + splitFilename[1]

    // check if components directory exists
    if componentsPath == "" {
        componentsPath = "./components"
    }
    if dirExists(componentsPath) {
        components.Init(componentsPath)
    }

    // read proxima file
    content, err := os.ReadFile(name + extension)
    if err != nil {
        exitOnError(err.Error())
    }

    out := os.Stdout

    // parse proxima file
    p := parser.New(string(content), name + extension)
    document := p.Parse()

    if len(p.Errors) != 0 {
        for _, err := range p.Errors {
            io.WriteString(out, err.String() + "\n")
        }
        return
    }

    // evaluate proxima file
    ev := evaluator.New(name + extension)
    evaluated := ev.Eval(document)
    if len(ev.Errors) != 0 {
        for _, err := range ev.Errors {
            io.WriteString(out, err.String() + "\n")
        }
        return
    }

    // generate html file
    html := PRE_HEAD

    for strings.HasPrefix(evaluated, "<link") || strings.HasPrefix(evaluated, "<script") || strings.HasPrefix(evaluated, "<title") {
        if strings.HasPrefix(evaluated, "<link") {
            splitHTML := strings.SplitN(evaluated, ">", 2)
            html += "\t" + splitHTML[0] + ">\n"
            evaluated = splitHTML[1]
        } else {
            splitHTML := strings.SplitN(evaluated, ">", 3)
            
            html += "\t" + splitHTML[0] + ">" + splitHTML[1] + ">\n"
            evaluated = splitHTML[2]
        }
    }

    html += POST_HEAD + evaluated + POST_BODY
    
    var output string
    if outputFilename != "" {
        output = outputFilename
    } else {
        output = name + ".html"
    }

    file, err := os.Create(output)
    if err != nil {
        exitOnError(err.Error())
    }
    defer file.Close()

    _, err = file.WriteString(html)
    if err != nil {
        exitOnError(err.Error())
    }
    
    msg := "\x1b[32m -> Generated " + output + "\x1b[0m"
    fmt.Println(msg)
}
