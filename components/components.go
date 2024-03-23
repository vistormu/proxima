package components

import (
    "fmt"
    "os"
    "bytes"
    "os/exec"
    "strings"
    "proxima/object"
)

type ComponentFunction func(args []string) object.Object
var Components map[string]ComponentFunction

type Component struct {
    Name string
    Extension string
    Content string
    NArgs int
}


func Init(dir string) {
    componentList := getComponents(dir)

    Components = make(map[string]ComponentFunction)
    for _, component := range componentList {
        Components[component.Name] = createComponentFunction(component)
    } 
}

func getComponents(dir string) []*Component {
    files, err := os.ReadDir(dir)
    if err != nil {
        fmt.Println("Error reading components directory")
        os.Exit(1)
    }

    var componentList []*Component
    for _, file := range files {
        filename := file.Name()

        if file.IsDir() {
            newDir := dir + filename + "/"
            componentList = append(componentList, getComponents(newDir)...)
            continue
        }

        component := getComponent(dir, filename)
        if component != nil {
            componentList = append(componentList, component)
        }
    }

    return componentList
}

func getComponent(dir, filename string) *Component {
    name := strings.Split(filename, ".")[0]
    extension := strings.Split(filename, ".")[1]

    switch extension {
    case "html":
        return createHTMLComponent(dir, filename, name)
    case "py":
        return createPythonComponent(dir, filename, name)
    }

    return nil
}


func createHTMLComponent(dir, filename, name string) *Component {
        content, err := os.ReadFile(dir + filename)
        if err != nil {
            fmt.Println("Error reading " + filename + " file")
            os.Exit(1)
        }
        
        nArgs := 0
        for _, char := range content {
            if char == '@' {
                nArgs++
            }
        }

        return &Component{
            Name: name,
            Content: string(content),
            NArgs: nArgs,
            Extension: "html",
        }
}

func createPythonComponent(dir, filename, name string) *Component {
    content, err := os.ReadFile(dir + filename)
    if err != nil {
        fmt.Println("Error reading " + filename + " file")
        os.Exit(1)
    }

    return &Component{
        Name: name,
        Content: string(content),
        NArgs: 0,
        Extension: "py",
    }
}

func createComponentFunction(component *Component) ComponentFunction {
    switch component.Extension {
    case "html":
        return createHTMLComponentFunction(component)
    default:
        return createPythonComponentFunction(component)
    }
}

func createHTMLComponentFunction(component *Component) ComponentFunction {
    return func(args []string) object.Object {
        if len(args) != component.NArgs {
            return &object.Error{ Message: fmt.Sprintf("wrong number of arguments in %s. got=%d, want=%d", component.Name, len(args), component.NArgs) }
        }
        value := component.Content
        for _, arg := range args {
            value = strings.Replace(value, "@", arg, 1)
        }
        if value[len(value) - 1] == '\n' {
            value = value[:len(value) - 1]
        }
        return &object.String{ Value: value }
    }
}

func createPythonComponentFunction(component *Component) ComponentFunction {
    return func(args []string) object.Object {
        evaluated, err := executePythonFunction(component.Content, component.Name, args)
        if err != nil {
            return &object.Error{ Message: err.Error() }
        }
        if evaluated[len(evaluated) - 1] == '\n' {
            evaluated = evaluated[:len(evaluated) - 1]
        }
        return &object.String{ Value: evaluated }
    }
}

func executePythonFunction(userFunction string, functionName string, args []string) (string, error) {
    nArgs := len(args)
    functionArgs := ""
    for i := 0; i < nArgs; i++ {
        functionArgs += "sys.argv[" + fmt.Sprint(i + 1) + "], "
    }

    pythonScript := userFunction + `

import sys
if __name__ == "__main__":
    print(function(
` + functionArgs + `))`

    cmdArgs := append([]string{"-c", pythonScript}, args...)
    cmd := exec.Command("python3", cmdArgs...)

    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr

    // Execute the command
    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("Error executing python function %s\n %s", functionName, stderr.String())
    }

    return strings.TrimSpace(out.String()), nil
}
