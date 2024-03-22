package commands

const (
    MAIN_EXT = ".prox"
    VERSION = "0.2.0"
)

func Execute(args []string) {
    switch args[0] {
    case "version":
        version()
    case "generate":
        generate(args[1:])
    case "watch":
        watch()
    case "help":
        help(args[1:])
    default:
        helpError()
    }
}
