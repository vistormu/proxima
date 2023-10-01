package components

import (
    "proxima/ast"
    "proxima/object"
    "strconv"
)

type ComponentFunction func(args []string, tagType ast.TagType) object.Object

var Builtins = map[string]ComponentFunction{
    // Headings
    "h1": h1,
    "h2": h2,
    "h3": h3,

    // Text styles
    "bold": bold,
    "italic": italic,
    "strike": strike,
    "uline": underline,
    "mark": mark,
    "code": code,

    // Lists
    "list": list,

    // Links
    "url": url,
    "email": email,

    // Images
    "image": image,

    // Other
    "break": breakline,
    "line": line,

    //preamble
    "style": style,
    "script": script,
    "title": title,
}

// Headings
func h1(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.WRAPPING {
        return &object.Error{Message: "@h1 can only be used as a wrapping tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@h1 can only have one argument"}
    }
    value := "<h1>\n\t" + args[0] + "\n</h1>\n"
    return &object.String{Value: value}
}
func h2(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.WRAPPING {
        return &object.Error{Message: "@h2 can only be used as a wrapping tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@h2 can only have one argument"}
    }
    value := "<h2>\n\t" + args[0] + "\n</h2>\n"
    return &object.String{Value: value}
}
func h3(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.WRAPPING {
        return &object.Error{Message: "@h3 can only be used as a wrapping tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@h3 can only have one argument"}
    }
    value := "<h3>\n\t" + args[0] + "\n</h3>\n"
    return &object.String{Value: value}
}

// Text styles
func bold(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@bold can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@bold can only have one argument"}
    }
    value := "<b>" + args[0] + "</b>"
    return &object.String{Value: value}
}
func italic(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@italic can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@italic can only have one argument"}
    }
    value := "<i>" + args[0] + "</i>"
    return &object.String{Value: value}
}
func strike(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@striket can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@striket can only have one argument"}
    }
    value := "<s>" + args[0] + "</s>"
    return &object.String{Value: value}
}
func underline(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@uline can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@uline can only have one argument"}
    }
    value := "<u>" + args[0] + "</u>"
    return &object.String{Value: value}
}
func mark(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@mark can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@mark can only have one argument"}
    }
    value := "<mark>" + args[0] + "</mark>"
    return &object.String{Value: value}
}
func code(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@code can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@code can only have one argument"}
    }
    value := "<code>" + args[0] + "</code>"
    return &object.String{Value: value}
}

// Lists
func list(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@list can only be used as a bracketed tag"}
    }
    if len(args) < 1 {
        return &object.Error{Message: "@list must have at least one argument"}
    }
    value := "<ul>\n"
    for _, arg := range args {
        value += "\t<li>" + arg + "</li>\n"
    }
    value += "</ul>\n"
    return &object.String{Value: value}
}

// Links
func url(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@url can only be used as a bracketed tag"}
    }
    if len(args) > 2 {
        return &object.Error{Message: "@url can only have one or two arguments"}
    }
    var value string
    if len(args) == 1 {
        value = "<a href=\"" + args[0] + "\">" + args[0] + "</a>"
    } else {
        value = "<a href=\"" + args[0] + "\">" + args[1] + "</a>"
    }
    return &object.String{Value: value}
}
func email(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@email can only be used as a bracketed tag"}
    }
    if len(args) > 2 {
        return &object.Error{Message: "@email can only have one or two arguments"}
    }
    var value string
    if len(args) == 1 {
        value = "<a href=\"mailto:" + args[0] + "\">" + args[0] + "</a>"
    } else {
        value = "<a href=\"mailto:" + args[0] + "\">" + args[1] + "</a>"
    }
    return &object.String{Value: value}
}

// Images
func image(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@image can only be used as a bracketed tag"}
    }
    if len(args) > 2 {
        return &object.Error{Message: "@image can only have one or two arguments"}
    }
    var value string
    if len(args) == 1 {
        value = "<img src=\"" + args[0] + "\">"
    } else {
        width, err := strconv.ParseFloat(args[1], 8)
        if err != nil {
            return &object.Error{Message: "@image takes a number as its second argument"}
        } else if width <= 0 || width > 1.0 {
            return &object.Error{Message: "@image can only have a width between 0 and 1"}
        } else {
            // parse width as a string
            widthString := strconv.FormatFloat(width * 720, 'f', -1, 64)
            value = "<img src=\"" + args[0] + "\" width=\"" + widthString + "\">"
        }
    }
    return &object.String{Value: value}
}

// Other
func breakline(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.SELF_CLOSING {
        return &object.Error{Message: "@break can only be used as a self-closing tag"}
    }
    if len(args) != 0 {
        return &object.Error{Message: "@break can't have any arguments"}
    }
    return &object.String{Value: "<br>\n"}
}
func line(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.SELF_CLOSING {
        return &object.Error{Message: "@line can only be used as a self-closing tag"}
    }
    if len(args) != 0 {
        return &object.Error{Message: "@line can't have any arguments"}
    }
    return &object.String{Value: "<hr>\n"}
}

// Preamble
func style(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@style can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@style can only have one argument"}
    }
    value := "<link rel=\"stylesheet\" href=\"" + args[0] + "\">"
    return &object.String{Value: value}
}
func script(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@script can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@script can only have one argument"}
    }
    value := "<script src=\"" + args[0] + "\"></script>"
    return &object.String{Value: value}
}
func title(args []string, tagType ast.TagType) object.Object {
    if tagType != ast.BRACKETED {
        return &object.Error{Message: "@title can only be used as a bracketed tag"}
    }
    if len(args) != 1 {
        return &object.Error{Message: "@title can only have one argument"}
    }
    value := "<title>" + args[0] + "</title>"
    return &object.String{Value: value}
}
