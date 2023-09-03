package builtins

type BuiltInFunction func(arg string) string

var Builtins = map[string]BuiltInFunction{
    // Alignment
    "center": center,
    "left": left,
    "right": right,

    // Headings
    "h1": h1,

    // Text styles
    "bold": bold,
    "italic": italic,
    "strike": strike,
    "underline": underline,

    // Lists

    // Links
    "url": url,

    // Images

    // Other
    "break": breakline,
}

// Alignment
func center(arg string) string {
    return "<div class=\"paragraph center\">\n\t" + arg + "\n</div>\n"
}
func left(arg string) string {
    return "<div class=\"paragraph left\">\n\t" + arg + "\n</div>\n"
}
func right(arg string) string {
    return "<div class=\"paragraph right\">\n\t" + arg + "\n</div>\n"
}

// Headings
func h1(arg string) string {
    return "<div class=\"h1\">\n\t" + arg + "\n</div>\n"
}

// Text styles
func bold(arg string) string {
    return `<b>` + arg + `</b>`
}
func italic(arg string) string {
    return `<i>` + arg + `</i>`
}
func strike(arg string) string {
    return `<s>` + arg + `</s>`
}
func underline(arg string) string {
    return `<u>` + arg + `</u>`
}

// Lists

// Links
func url(arg string) string {
    return `<a href="` + arg + `">` + arg + `</a>`
}

// Images

// Other
func breakline(arg string) string {
    return `<br>`
}
