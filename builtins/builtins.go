package builtins

type BuiltInFunction func(arg string) string

var Builtins = map[string]BuiltInFunction{
    // Alignment
    "center": center,
    "left": left,
    "right": right,

    // Headings
    "h0": h0,
    "h1": h1,

    // Text styles
    "bold": bold,
    "italic": italic,
    "striket": strike,
    "uline": underline,

    "monospace": monospace,

    // Lists

    // Links
    "url": url,

    // Images

    // Other
    "break": breakline,
    "line": line,
}

// Alignment
func center(arg string) string {
    return "<div class=\"paragraph\" id=\"center\">\n\t" + arg + "\n</div>\n"
}
func left(arg string) string {
    return "<div class=\"paragraph\" id=\"left\">\n\t" + arg + "\n</div>\n"
}
func right(arg string) string {
    return "<div class=\"paragraph\" id=\"right\">\n\t" + arg + "\n</div>\n"
}

// Headings
func h0(arg string) string {
    return "<div class=\"h0\">\n\t" + arg + "\n</div>\n"
}
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

func monospace(arg string) string {
    return `<div class="paragraph" id="monospace">` + arg + `</div>`
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
func line(arg string) string {
    return `<hr>`
}
