package main

import (
    "fmt"
    "os"
    "strings"
    "time"
    "regexp"
    "path/filepath"
    "proxima/parser"
    "proxima/evaluator"
    "proxima/components"
)

const (
    MAIN_EXT = ".prox"
    VERSION = "0.2.0"
)

func exitOnError(msg string) {
    fmt.Println("\x1b[31m-> |build| " + msg + "\x1b[0m")
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
}

func generate(filename string, componentsPath string) {
    before := time.Now()

    // check if components directory exists
    if componentsPath == "" {
        componentsPath = "./components"
    }
    if dirExists(componentsPath) {
        components.Init(componentsPath)
    }

    // read proxima file
    content, err := os.ReadFile(filename)
    if err != nil {
        exitOnError(err.Error())
    }

    // parse proxima file
    p := parser.New(string(content), filename)
    document := p.Parse()
    if len(p.Errors) != 0 {
        for _, err := range p.Errors {
            fmt.Println(err.String())
        }
        os.Exit(1)
    }

    // evaluate proxima file
    ev := evaluator.New(filename)
    evaluated := ev.Eval(document)
    if len(ev.Errors) != 0 {
        for _, err := range ev.Errors {
            fmt.Println(err.String())
        }
        os.Exit(1)
    }

    // find elements that should be contained in the head tag
    headHtml := ""
    bodySplit := strings.Split(formatHTML(evaluated), "\n")

    for i, line := range bodySplit {
        found := false
        for _, element := range []string{"title", "meta", "link", "style", "script"} {
            if strings.HasPrefix(line, "<" + element) || strings.HasPrefix(line, "</" + element) {
                headHtml += line + "\n"
                bodySplit[i] = ""
                found = true
            }
        }
        if !found {
            break
        }
    }

    bodyHtml := strings.TrimSpace(strings.Join(bodySplit, "\n"))

    // generate html file
    preHead := "<!DOCTYPE html>\n<html>\n<head>\n"
    postHead := "</head>\n<body>\n"
    postBody := "</body>\n</html>"
    html := preHead + headHtml + postHead + bodyHtml + postBody
    
    output := strings.TrimSuffix(filename, MAIN_EXT) + ".html"
    file, err := os.Create(output)
    if err != nil {
        exitOnError(err.Error())
    }
    defer file.Close()

    _, err = file.WriteString(html)
    if err != nil {
        exitOnError(err.Error())
    }

    after := time.Now()
    elapsed := after.Sub(before)
    elapsed = elapsed - elapsed%time.Millisecond

    msg := fmt.Sprintf("\x1b[32m-> Generated %s (%s)\x1b[0m", output, elapsed)
    fmt.Println(msg)
}

func formatHTML(html string) string {
	// Step 1: Insert new lines before "<", except at the beginning
	regexNewLine := regexp.MustCompile(`(?m)(<)`)
	formatted := regexNewLine.ReplaceAllString(html, "\n$1")

	// Step 2: Trim leading whitespace from each line
	lines := strings.Split(formatted, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	formatted = strings.Join(lines, "\n")

	// Step 3: Replace multiple newlines with a single newline
	regexMultiNewLine := regexp.MustCompile(`\n+`)
	formatted = regexMultiNewLine.ReplaceAllString(formatted, "\n")

	return strings.TrimSpace(formatted)
}

