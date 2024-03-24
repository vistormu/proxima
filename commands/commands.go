package commands

const (
    MAIN_EXT = ".prox"
    VERSION = "0.2.0"
)

func Execute(args []string) {
    if len(args) < 2 {
        helpError()
    }
    args = args[1:]

    switch args[0] {
    case "version":
        version()
    case "generate":
        err := generate(args[1:])
        if err != nil {
            exitOnError(err.Error())
        }
    case "watch":
        err := watch(args[1:])
        if err != nil {
            exitOnError(err.Error())
        }
    case "help":
        help(args[1:])
    default:
        helpError()
    }
}
