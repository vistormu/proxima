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

func parseArgs(args []string) (map[string]string, error) {
    equivalentFlags := map[string]string{
        "-o":           "--output",
        "--output":     "--output",
    }

    parsedFlags := map[string]string{
        "--output":    "",
    }

    // parse flags
    for i := 1; i < len(args); i++ {
        if !strings.HasPrefix(args[i], "-") {
            continue
        }

        flag, ok := equivalentFlags[args[i]]
        if !ok {
            return nil, errors.NewCliError(errors.UNKNOWN_FLAG, args[i])
        }

        if i+1 >= len(args) {
            return nil, errors.NewCliError(errors.MISSING_FLAG_VALUE, args[i])
        }

        if strings.HasPrefix(args[i+1], "-") {
            return nil, errors.NewCliError(errors.MISSING_FLAG_VALUE, args[i])
        }

        parsedFlags[flag] = args[i+1]
        i++
    }

    return parsedFlags, nil
}

func make_(args []string) error {
    if len(args) < 1 {
        return errors.NewCliError(errors.WRONG_N_ARGS, 1, len(args))
    }

    // get file
    filename := args[0]
    if !strings.HasSuffix(filename, MAIN_EXT) {
        return errors.NewCliError(errors.INVALID_FILE_EXTENSION, filename)
    }

    // parse flags
    flags, err := parseArgs(args)
    if err != nil {
        return err
    }

    if flags["--output"] == "" {
        return errors.NewCliError(errors.NO_OUTPUT_FLAG)
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
        return err
    }

    // tokenize
    t := tokenizer.New([]rune(string(content)))
    tokens := t.Tokenize()

    // parse
    p := parser.New(tokens, filename, c)
    expressions := p.Parse()

    if len(p.Errors) > 0 {
        for _, e := range p.Errors {
            fmt.Println(e.Error())
        }
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
        return err
    }

    fmt.Printf("\x1b[32;1m\"%s\" generated successfully! (%d ms)\x1b[0m\n\n", flags["--output"], time.Since(begin).Milliseconds())

    return nil
}
