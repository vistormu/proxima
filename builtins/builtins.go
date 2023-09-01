package builtins

type BuiltInFunction func(args... string) string

var Builtins = map[string]BuiltInFunction{
    "h1": h1,
    "bold": bold,
    "italic": italic,
}

func h1(args... string) string {
    return `<h1>` + args[0] + `</h1>`
}
func bold(args... string) string {
    return `<b>` + args[0] + `</b>`
}
func italic(args... string) string {
    return `<i>` + args[0] + `</i>`
}
