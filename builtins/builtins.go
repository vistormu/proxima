package builtins

type BuiltInFunction func(arg string) string

var Builtins = map[string]BuiltInFunction{
    // Alignment
    "center": center,
    "justify": justify,
    "left": left,
    "right": right,

    // Headings
    "h1": h1,
    "h2": h2,
    "h3": h3,

    // Text styles
    "bold": bold,
    "italic": italic,
    "strike": strike,
    "underline": underline,

    // Lists
    "bulletlist": bulletlist,
    "numberlist": numberlist,

    // Links
    "url": url,
    "email": email,

    // Images
    "image": image,

    // Other
    "break": breakline,
}

// Alignment
func center(arg string) string {
    return `<div style="text-align: center">` + arg + `</div>`
}
func justify(arg string) string {
    return `<div style="text-align: justify">` + arg + `</div>`
}
func left(arg string) string {
    return `<div style="text-align: left">` + arg + `</div>`
}
func right(arg string) string {
    return `<div style="text-align: right">` + arg + `</div>`
}

// Headings
func h1(arg string) string {
    return `<h1>` + arg + `</h1>`
}
func h2(arg string) string {
    return `<h2>` + arg + `</h2>`
}
func h3(arg string) string {
    return `<h3>` + arg + `</h3>`
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
func bulletlist(arg string) string {
    return `<ul>` + arg + `</ul>`
}
func numberlist(arg string) string {
    return `<ol>` + arg + `</ol>`
}

// Links
func url(arg string) string {
    return `<a href="` + arg + `">` + arg + `</a>`
}
func email(arg string) string {
    return `<a href="mailto:` + arg + `">` + arg + `</a>`
}

// Images
func image(arg string) string {
    return `<img src="` + arg + `">`
}

// Other
func breakline(arg string) string {
    return `<br>`
}
