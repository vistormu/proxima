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

func processDirectory(dirPath string) error {
    files, err := os.ReadDir(dirPath)
    if err != nil {
        return err
    }

    for _, file := range files {
        fullPath := filepath.Join(dirPath, file.Name())
        if file.IsDir() {
            // If it's a directory, recurse into it
            err := processDirectory(fullPath)
            if err != nil {
                return err // Propagate errors
            }
        } else if strings.HasSuffix(file.Name(), MAIN_EXT) {
            // If it's a file with the correct extension, apply 'generate'
            generate(fullPath)
        }
    }

    return nil
}

func main() {
    if len(os.Args) != 2 {
        exitOnError("usage: proxima <filename>.prox or proxima all to process all .prox files")
    }

    if os.Args[1] == "all" {
        error := processDirectory("./")
        if error != nil {
            exitOnError("error.Error()")
        }
    } else {
        generate(os.Args[1])
    }

}

func generate(filename string) {
    // <name>.prox
    splitFilename := strings.SplitN(filename, ".", 2)
    name := splitFilename[0]
    extension := "." + splitFilename[1]
    if extension != MAIN_EXT {
        exitOnError("file must have .prox extension")
    }

    // check if components directory exists
    if dirExists("./components") {
        components.Init()
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

    file, err := os.Create(name + ".html")
    if err != nil {
        exitOnError(err.Error())
    }
    defer file.Close()

    _, err = file.WriteString(html)
    if err != nil {
        exitOnError(err.Error())
    }
    
    msg := "\x1b[32m -> Generated " + name + ".html\x1b[0m"
    fmt.Println(msg)
}
