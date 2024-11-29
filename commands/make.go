package commands

import (
    "fmt"
    "os"
    "strings"
    "time"

    "proxima/config"
    "proxima/evaluator"
    "proxima/parser"
    "proxima/tokenizer"
    "proxima/errors"
)

var flags = map[string]string{
    "-o": "--output",
    "--output": "--output",
    "-output": "--output",
}

func parseArgs(args []string) (map[string]string, error) {
    // parse flags
    var parsedFlags = make(map[string]string)
    for i := 1; i < len(args); i++ {
        if !strings.HasPrefix(args[i], "-") {
            continue
        }

        flag, ok := flags[args[i]]
        if !ok {
            closestMatch := findClosestMatch(args[i], keys(flags))
            return nil, errors.New(errors.FLAG, args[i], closestMatch)
        }

        if i+1 >= len(args) {
            return nil, errors.New(errors.FLAG_VALUE, args[i])
        }

        if strings.HasPrefix(args[i+1], "-") {
            return nil, errors.New(errors.FLAG_VALUE, args[i])
        }

        parsedFlags[flag] = args[i+1]
        i++
    }

    return parsedFlags, nil
}

func make_(args []string) error {
    if len(args) < 1 {
        return errors.New(errors.N_ARGS, 1, len(args))
    }

    // get file
    filename := args[0]
    if !strings.HasSuffix(filename, MAIN_EXT) {
        wrongExtension := strings.Split(filename, ".")[1]
        return errors.New(errors.EXTENSION, "." + wrongExtension)
    }

    // parse flags
    flags, err := parseArgs(args)
    if err != nil {
        return err
    }

    if flags["--output"] == "" {
        return errors.New(errors.OUTPUT_FLAG)
    }

    // load config
    c, err := config.LoadConfig()
    if err != nil {
        return err
    }

    begin := time.Now()

    // read file
    content, err := os.ReadFile(filename)
    if err != nil {
        return errors.New(errors.READ_FILE, filename)
    }

    // tokenize
    t := tokenizer.New([]rune(string(content)))
    tokens := t.Tokenize()

    // parse
    p := parser.New(tokens, filename, c)
    expressions := p.Parse()

    if len(p.Errors) > 0 {
        errorMsg := ""
        for _, e := range p.Errors {
            errorMsg += e.Error() + "\n"
        }
        return fmt.Errorf("%s", errorMsg)
    }

    // evaluate
    e, err := evaluator.New(expressions, filename, c)
    if err != nil {
        return err
    }
    defer e.Close()

    result, err := e.Evaluate()
    if err != nil {
        return err
    }

    // save result as output file
    err = os.WriteFile(flags["--output"], []byte(result), 0644)
    if err != nil {
        return errors.New(errors.CREATE_FILE, flags["--output"])
    }

    fmt.Printf("\x1b[32;1m\"%s\" generated successfully! (%d ms)\x1b[0m\n\n", flags["--output"], time.Since(begin).Milliseconds())

    return nil
}
