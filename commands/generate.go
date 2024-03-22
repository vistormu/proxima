package commands

import (
    "fmt"
    "os"
    "regexp"
    "strings"
    "time"
    "proxima/components"
    "proxima/parser"
    "proxima/evaluator"
)

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
    postBody := "\n</body>\n</html>"
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

