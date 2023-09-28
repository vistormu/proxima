package components

import (
    "fmt"
    "os"
    "strings"
    "proxima/ast"
    "proxima/object"
)

var Components map[string]ComponentFunction

type Component struct {
    Name string
    Content string
    TagType ast.TagType
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
        tagTypeString := strings.Split(name, "-")[1]
        name = strings.Split(name, "-")[0]
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
        
        var tagType ast.TagType
        switch tagTypeString {
        case "s":
            tagType = ast.SELF_CLOSING
        case "b":
            tagType = ast.BRACKETED
        case "w":
            tagType = ast.WRAPPING
        default:
            fmt.Println("Error: invalid tag type in component file name")
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
            TagType: tagType,
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
    switch component.TagType {
    case ast.SELF_CLOSING:
        return func(args []string, tagType ast.TagType) object.Object {
            if len(args) != component.NArgs {
                return &object.Error{ Message: fmt.Sprintf("wrong number of arguments in %s. got=%d, want=%d", component.Name, len(args), component.NArgs) }
            }
            return &object.String{ Value: component.Content }
        }

    case ast.BRACKETED:
        return func(args []string, tagType ast.TagType) object.Object { 
            if len(args) != component.NArgs {
                return &object.Error{ Message: fmt.Sprintf("wrong number of arguments in %s. got=%d, want=%d", component.Name, len(args), component.NArgs) }
            }
            value := component.Content
            for _, arg := range args {
                value = strings.Replace(value, "@", arg, 1)
            }
            return &object.String{ Value: value }
        }

    case ast.WRAPPING:
        return func(args []string, tagType ast.TagType) object.Object {
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

    return nil
}
