package components

import (
    "fmt"
    "os"
    "strings"
    "proxima/object"
)

var Components map[string]ComponentFunction

type Component struct {
    Name string
    Content string
    NArgs int
}


func Init() {
    files, err := os.ReadDir("./components")
    if err != nil {
        fmt.Println("Error reading components directory")
        os.Exit(1)
    }

    componentList := make([]*Component, len(files))
    for i, file := range files {
        filename := file.Name()
        name := strings.Split(filename, ".")[0]
        extension := strings.Split(filename, ".")[1]

        if extension != "html" {
            fmt.Println("Error: component file must have .html extension")
            os.Exit(1)
        }

        content, err := os.ReadFile("./components/" + filename)
        if err != nil {
            fmt.Println("Error reading" + filename + "file")
            os.Exit(1)
        }
        
        nArgs := 0
        for _, char := range content {
            if char == '@' {
                nArgs++
            }
        }

        componentList[i] = &Component{
            Name: name,
            Content: string(content),
            NArgs: nArgs,
        }
    } 
    Components = make(map[string]ComponentFunction)
    fillMap(componentList)
}

func fillMap(componentList []*Component) {
    for _, component := range componentList {
        function := createFunction(component)
        if function != nil {
            Components[component.Name] = function
        }
    } 
}

func createFunction(component *Component) ComponentFunction {
    switch component.NArgs {
    case 0:
        return func(args []string) object.Object {
            if len(args) != component.NArgs {
                return &object.Error{ Message: fmt.Sprintf("wrong number of arguments in %s. got=%d, want=%d", component.Name, len(args), component.NArgs) }
            }
            return &object.String{ Value: component.Content }
        }

    default:
        return func(args []string) object.Object { 
            if len(args) != component.NArgs {
                return &object.Error{ Message: fmt.Sprintf("wrong number of arguments in %s. got=%d, want=%d", component.Name, len(args), component.NArgs) }
            }
            value := component.Content
            for _, arg := range args {
                value = strings.Replace(value, "@", arg, 1)
            }
            return &object.String{ Value: value }
        }
    }
}
