package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vistormu/go-dsa/ansi"
)

// VERSION is set at release time by GoReleaser (-ldflags -X proxima/cmd.VERSION=...).
var VERSION = "dev"

const (
	MAIN_EXT = ".prox"
	NAME     = ansi.Italic + ansi.Magenta + "proxima" + ansi.Reset
)

var rootCmd = &cobra.Command{
	Use:     "proxima",
	Version: VERSION,
}

func init() {
	rootCmd.SetHelpTemplate(defaultHelp)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
