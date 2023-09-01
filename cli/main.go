package main

import (
    "io"
    "io/ioutil"
    "os"
    "strings"
    "proxima/parser"
    "proxima/evaluator"
)

const (
    MAIN_EXT = ".prox"
    STYLES_EXT = ".prox"
)

func main() {
    if len(os.Args) != 2 {
        panic("Usage: proxima <filename>")
    }

    filename := os.Args[1]
    if !strings.HasSuffix(filename, MAIN_EXT) {
        panic("File must have .prox extension")
    }

    content, err := ioutil.ReadFile(filename)
    if err != nil {
        panic(err)
    }

    out := os.Stdout
    parser := parser.New(string(content))
    document := parser.Parse()

    if len(parser.Errors) != 0 {
        for _, err := range parser.Errors {
            io.WriteString(out, err.String() + "\n")
        }
        return
    }

    evaluated := evaluator.Eval(document)

    file, err := os.Create("index.html")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    _, err = file.WriteString(evaluated)
    if err != nil {
        panic(err)
    }
}
