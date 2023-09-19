package main

import (
    "io"
    "io/ioutil"
    "os"
    "os/exec"
    "strings"
    "proxima/parser"
    "proxima/evaluator"
)

const (
    MAIN_EXT = ".prox"
)

func main() {
    if len(os.Args) != 2 && len(os.Args) != 3 {
        panic("Usage: proxima <filename>")
    }

    // <filename>.prox
    filename := os.Args[1][:strings.LastIndex(os.Args[1], ".")]
    extension := os.Args[1][strings.LastIndex(os.Args[1], "."):]
    if extension != MAIN_EXT {
        panic("File must have .prox extension")
    }

    content, err := ioutil.ReadFile(filename + extension)
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

    cmdPrompt := []string{"wkhtmltopdf", "-R", "25mm", "-B", "25mm", "-L", "25mm", "-T", "25mm", "--enable-local-file-access", "index.html", filename + ".pdf"}
    cmd := exec.Command(cmdPrompt[0], cmdPrompt[1:]...)
    err = cmd.Run()
    if err != nil {
        panic(err)
    }

    htmlFlag := false
    for _, arg := range os.Args {
        if arg == "--html" {
            htmlFlag = true
            break
        }
    }
    
    if !htmlFlag {
        cmd = exec.Command("rm", "index.html")
        err = cmd.Run()
        if err != nil {
            panic(err)
        }
    }

    cmd = exec.Command("open", filename + ".pdf")
    err = cmd.Run()
    if err != nil {
        panic(err)
    }
}
