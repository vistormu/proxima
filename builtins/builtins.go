package builtins

import "proxima/ast"

type BuiltInFunction func(content string, tagType ast.TagType) string

var Builtins = map[string]BuiltInFunction{
    // Alignment
    "center": center,
    "left": left,
    "right": right,

    // Headings
    "h1": h1,
    "h2": h2,
    "h3": h3,

    // Text styles
    "bold": bold,
    "italic": italic,
    "striket": strike,
    "uline": underline,

    // Lists

    // Links
    "url": url,

    // Images
    "image": image,

    // Other
    "break": breakline,
    "line": line,
}

// Alignment
func center(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.WRAPPING {
        result = "<div class=\"paragraph\" id=\"center\">\n\t" + arg + "\n</div>\n"
    }
    return result
}
func left(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.WRAPPING {
        result = "<div class=\"paragraph\" id=\"left\">\n\t" + arg + "\n</div>\n"
    }
    return result
}
func right(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.WRAPPING {
        result = "<div class=\"paragraph\" id=\"right\">\n\t" + arg + "\n</div>\n"
    }
    return result
}

// Headings
func h1(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.WRAPPING {
        result = "<div class=\"h1\">\n\t" + arg + "\n</div>\n"
    }
    return result
}
func h2(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.WRAPPING {
        result = "<div class=\"h2\">\n\t" + arg + "\n</div>\n"
    }
    return result
}
func h3(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.WRAPPING {
        result = "<div class=\"h3\">\n\t" + arg + "\n</div>\n"
    }
    return result
}

// Text styles
func bold(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.BRACKETED {
        result = "<b>" + arg + "</b>"
    }
    return result
}
func italic(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.BRACKETED {
        result = "<i>" + arg + "</i>"
    }
    return result
}
func strike(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.BRACKETED {
        result = "<s>" + arg + "</s>"
    }
    return result
}
func underline(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.BRACKETED {
        result = "<u>" + arg + "</u>"
    }
    return result
}

// Lists

// Links
func url(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.BRACKETED {
        result = "<a href=\"" + arg + "\">" + arg + "</a>"
    }
    return result
}

// Images
func image(arg string, tagTYpe ast.TagType) string {
    result := ""
    if tagTYpe == ast.BRACKETED {
        result = "<img src=\"" + arg + "\">"
    }
    return result
}

// Other
func breakline(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.SELF_CLOSING {
        result = "<br>"
    }
    return result
}
func line(arg string, tagType ast.TagType) string {
    result := ""
    if tagType == ast.SELF_CLOSING {
        result = "<hr>"
    }
    return result
}
