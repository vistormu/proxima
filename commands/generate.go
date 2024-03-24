package commands

import (
    "fmt"
    "os"
    "strings"
    "time"
    "proxima/components"
    "proxima/parser"
    "proxima/evaluator"
)

func generate(args []string) error {
    // flags
    componentsPath, err := getComponentsPath(args)
    if err != nil {
        return err
    }
    isRecursive := getIsRecursive(args)

    filesAndDirs := []string{}
    for _, arg := range args {
        if arg != "" {
            filesAndDirs = append(filesAndDirs, arg)
        }
    }
    
    // errors
    if len(filesAndDirs) == 0 {
        return fmt.Errorf("no files specified\nCheck 'proxima help generate' for more information")
    }
    for _, arg := range filesAndDirs {
        isDirectory, _ := isDir(arg)
        if isDirectory {
            directoryExists, _ := dirExists(arg)
            if !directoryExists {
                return fmt.Errorf("directory %s does not exist", arg)
            }
        } else {
            if !strings.HasSuffix(arg, MAIN_EXT) {
                return fmt.Errorf("file %s is not a .prox file", arg)
            }
        }
    }

    // generate
    for _, arg := range filesAndDirs {
        isDirectory, _ := isDir(arg)
        if isDirectory {
            err := generateDir(arg, componentsPath, isRecursive)
            if err != nil {
                return err
            }
        } else {
            err := generateFile(arg, componentsPath)
            if err != nil {
                return err
            }
        }
    }

    return nil
}


// =====
// FLAGS
// =====
func getComponentsPath(args []string) (string, error) {
    componentsPath := ""
    for i, arg := range args {
        if arg == "-c" && i + 1 < len(args) {
            componentsPath = args[i + 1]
            args[i] = ""
            args[i + 1] = ""
            break
        }
    }

    componentsPathExists, err := dirExists(componentsPath)
    if err != nil {
        return "", err
    }

    if componentsPath != "" && !componentsPathExists {
        return "", fmt.Errorf("components directory does not exist")
    }

    if componentsPath != "" && !strings.HasSuffix(componentsPath, "/") {
        componentsPath += "/"
    }

    defaultComponentsPathExists, _ := dirExists("./components")
    if componentsPath == "" && defaultComponentsPathExists {
        componentsPath = "./components/"
    }

    return componentsPath, nil
}
func getIsRecursive(args []string) bool {
    for i, arg := range args {
        if arg == "-r" {
            args[i] = ""
            return true
        }
    }
    return false
}

// ========
// GENERATE
// ========
func generateDir(dir string, componentsPath string, isRecursive bool) error {
    if dir == "." {
        dir = "./"
    } else if !strings.HasSuffix(dir, "/") {
        dir += "/"
    }

    var files []string
    var err error
    if isRecursive {
        files, err = getAllFiles(dir)
        if err != nil {
            return err
        }
    } else {
        files, err = getFilesInDir(dir)
        if err != nil {
            return err
        }
    }
    for _, file := range files {
        err = generateFile(file, componentsPath)
        if err != nil {
            return err
        }
    }

    return nil
}

func generateFile(filename string, componentsPath string) error {
    before := time.Now()

    // initialize components
    if componentsPath != "" {
        components.Init(componentsPath)
    }

    // read
    content, err := os.ReadFile(filename)
    if err != nil {
        return err
    }

    // parse
    p := parser.New(string(content), filename)
    document := p.Parse()
    if len(p.Errors) != 0 {
        errorString := ""
        for _, err := range p.Errors {
            errorString += err.String() + "\n"
        }
        return fmt.Errorf(errorString)
    }

    // evaluate
    ev := evaluator.New(filename)
    evaluated := ev.Eval(document)
    if len(ev.Errors) != 0 {
        errorString := ""
        for _, err := range ev.Errors {
            errorString += err.String() + "\n"
        }
        return fmt.Errorf(errorString)
    }

    // format
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

    preHead := "<!DOCTYPE html>\n<html>\n<head>\n"
    postHead := "</head>\n<body>\n"
    postBody := "\n</body>\n</html>"
    html := preHead + headHtml + postHead + bodyHtml + postBody
    
    // generate html file
    output := strings.TrimSuffix(filename, MAIN_EXT) + ".html"
    file, err := os.Create(output)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(html)
    if err != nil {
        return err
    }

    after := time.Now()
    elapsed := after.Sub(before)
    elapsed = elapsed - elapsed%time.Millisecond

    msg := fmt.Sprintf("\x1b[32m-> Generated %s (%s)\x1b[0m", output, elapsed)
    fmt.Println(msg)

    return nil
}

func formatHTML(html string) string {
	var sb strings.Builder

	inTag := false
	for i := 0; i < len(html); i++ {
		if html[i] == '<' && !(i+1 < len(html) && html[i+1] == '/') {
			if i > 0 {
				sb.WriteRune('\n')
			}
		}
		sb.WriteByte(html[i])

		if html[i] == '<' {
			inTag = true
		} else if inTag && html[i] == '>' {
			inTag = false
		}
	}

	formatted := sb.String()
	formatted = strings.TrimSpace(formatted)
	formatted = strings.ReplaceAll(formatted, "\n\n", "\n")

	return formatted
}
