package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"proxima/internal/config"
	"proxima/internal/evaluator"
	"proxima/internal/parser"
	"proxima/internal/tokenizer"

	"github.com/spf13/cobra"
	"github.com/vistormu/go-dsa/errors"
)

var outputFilename string

var makeCmd = &cobra.Command{
	Use:   "make <file.prox> [flags]",
	Short: "transpile the proxima file to an output file",
	Long:  `the make command transpiles the proxima file to an output file`,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// get file
		filename := args[0]
		if !strings.HasSuffix(filename, MAIN_EXT) {
			wrongExtension := strings.Split(filename, ".")[1]
			return errors.New(WrongExtension).With(
				"expected", MAIN_EXT,
				"got", wrongExtension,
			)
		}
		// load config
		c, err := config.LoadConfig("proxima.toml")
		if err != nil {
			return err
		}

		begin := time.Now()

		// read file
		content, err := os.ReadFile(filename)
		if err != nil {
			return errors.New(ReadFile).With(
				"filename", filename,
			).Wrap(err)
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
		err = os.WriteFile(outputFilename, []byte(result), 0644)
		if err != nil {
			return errors.New(WriteFile).With(
				"filename", outputFilename,
			).Wrap(err)
		}

		successMsg := "\x1b[32m"
		successMsg += "-> proxima finished successfully!\x1b[0m\n"
		successMsg += fmt.Sprintf("   |> output file: %s\n", outputFilename)
		successMsg += fmt.Sprintf("   |> time elapsed: %d ms\n", time.Since(begin).Milliseconds())
		fmt.Println(successMsg)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)
	makeCmd.Flags().StringVarP(&outputFilename, "output", "o", "", "output filename")
}
