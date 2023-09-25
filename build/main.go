package main

import (
    "fmt"
    "io"
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

    content, err := os.ReadFile(filename + extension)
    if err != nil {
        fmt.Println("Error reading file")
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

    file, err := os.Create(filename + ".html")
    if err != nil {
        fmt.Println("Error creating file")
        panic(err)
    }
    defer file.Close()

    _, err = file.WriteString(evaluated)
    if err != nil {
        fmt.Println("Error writing to file")
        panic(err)
    }

    htmlFlag := false
    for _, arg := range os.Args {
        if arg == "--html" {
            htmlFlag = true
            break
        }
    }

    if htmlFlag {
        fmt.Println("HTML file generated")
        return
    }

    cmdPrompt := []string{"wkhtmltopdf", "-R", "25mm", "-B", "25mm", "-L", "25mm", "-T", "25mm", "--enable-local-file-access", filename + ".html", filename + ".pdf"}
    cmd := exec.Command(cmdPrompt[0], cmdPrompt[1:]...)
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error running wkhtmltopdf")
        panic(err)
    }

    cmd = exec.Command("rm", filename + ".html")
    err = cmd.Run()
    if err != nil {
        fmt.Println("Error removing index.html")
        panic(err)
    }
}
