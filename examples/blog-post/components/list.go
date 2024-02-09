package components

func list(items []string) string {
    value := "<ul>\n"
    for _, item := range items {
        value += "\t<li>" + item + "</li>\n"
    }
    value += "</ul>\n"
    return value
}

