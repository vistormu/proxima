package components

import (
    "proxima/object"
)


var Builtins = map[string]ComponentFunction{
    "style": style,
    "script": script,
    "title": title,
}

func style(args []string) object.Object {
    if len(args) != 1 {
        return &object.Error{Message: "@style can only have one argument"}
    }
    value := "<link rel=\"stylesheet\" href=\"" + args[0] + "\">"
    return &object.String{Value: value}
}
func script(args []string) object.Object {
    if len(args) != 1 {
        return &object.Error{Message: "@script can only have one argument"}
    }
    value := "<script src=\"" + args[0] + "\"></script>"
    return &object.String{Value: value}
}
func title(args []string) object.Object {
    if len(args) != 1 {
        return &object.Error{Message: "@title can only have one argument"}
    }
    value := "<title>" + args[0] + "</title>"
    return &object.String{Value: value}
}
