package commands

import (
    "fmt"
    "proxima/errors"
)

var helpFunctions = map[string]func(){
    "init": helpInit,
    "make": helpMake,
    "version": helpVersion,
    "help": helpHelp,
}

func help(args []string) error {
    if len(args) > 1 {
        return errors.New(errors.N_ARGS, 1, len(args))
    }

    if len(args) == 0 {
        helpBasic()
        return nil
    }

    function, ok := helpFunctions[args[0]]
    if !ok {
        closestMatch := findClosestMatch(args[0], keys(helpFunctions))
        return errors.New(errors.COMMAND, args[0], closestMatch)
    }

    function()

    return nil
}

func helpBasic() {
    msg := "\x1b[32musage\x1b[0m:\n" +
    "    proxima <command> [arguments]\n\n" +
    "\x1b[32mcommands\x1b[0m:\n" +
    "    \x1b[35minit\x1b[0m       Create a new proxima project\n" +
    "    \x1b[35mmake\x1b[0m       Generates the specified file from a .prox file\n" +
    "    \x1b[35mversion\x1b[0m    Display the current version\n\n" +
    "For more information on a specific command, use 'proxima help <command>'."

    fmt.Println("\x1b[35mproxima\x1b[0m " + string(VERSION) + "\n\n" + msg)
}

func helpMake() {
    msg := "\x1b[35mproxima\x1b[0m make\n\n" +
    "\x1b[32musage\x1b[0m:\n" +
    "    proxima make <file> -o <output_file>\n\n" +
    "\x1b[32mdescription\x1b[0m:\n" +
    "    generate the specified new file from a .prox file\n\n" +
    "\x1b[32mflags\x1b[0m:\n" +
    "    -o <output_file>        the output file name\n"

    fmt.Println(msg)
}

func helpVersion() {
    msg := "\x1b[35mproxima\x1b[0m version\n\n" +
    "\x1b[32musage\x1b[0m:\n" +
    "    proxima version\n\n" +
    "\x1b[32mdescription\x1b[0m:\n" +
    "    display the current version of proxima"

    fmt.Println(msg)
}

func helpHelp() {
    msg := "\x1b[35mproxima\x1b[0m help\n\n" +
    "\x1b[32musage\x1b[0m:\n" +
    "    proxima help [command]\n\n" +
    "\x1b[32mdescription\x1b[0m:\n" +
    "    display help information for proxima\n\n" +
    "why are you even reading this?"

    fmt.Println(msg)
}

func helpInit() {
    msg := "\x1b[35mproxima\x1b[0m init\n\n" +
    "\x1b[32musage\x1b[0m:\n" +
    "    proxima init\n\n" +
    "\x1b[32mdescription\x1b[0m:\n" +
    "    create a new proxima project. it will create a new directory with the following structure:\n" +
    "        - components/      the default directory for creating new components\n" +
    "        - config.toml      the default configuration file\n" +
    "        - main.prox        an example proxima file\n"

    fmt.Println(msg)
}
