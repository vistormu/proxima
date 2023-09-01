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
    p := parser.New(string(content))
    document := p.Parse()

    if len(p.Errors) != 0 {
        for _, err := range p.Errors {
            io.WriteString(out, err.String() + "\n")
        }
        return
    }

    ev := evaluator.New()
    evaluated := ev.Eval(document)
    if len(ev.Errors) != 0 {
        for _, err := range ev.Errors {
            io.WriteString(out, err.String() + "\n")
        }
        return
    }

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
