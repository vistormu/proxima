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
    STYLES_EXT = ".css"
)

const preamble = `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Document</title>

    <style>
        @page {
            size: A4;
            margin: 27mm 16mm 27mm 16mm;
        }
        .paragraph {
            margin-top: 20px;
            margin-bottom: 20px;
            text-indent: 20px;
            text-align: justify;
        }
        .h1 {
            font-size: 32px;
            font-weight: bold;
            font-family: sans-serif;
        }
        .center {
            text-align: center;
        }
        .right {
            text-align: right;
        }
    </style>
</head>

<body>
`
const postamble = `
</body>
</html>
`

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

    _, err = file.WriteString(preamble)
    if err != nil {
        panic(err)
    }

    _, err = file.WriteString(evaluated)
    if err != nil {
        panic(err)
    }

    _, err = file.WriteString(postamble)
    if err != nil {
        panic(err)
    }
}
