package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/vistormu/go-dsa/ansi"
)

const (
	MAIN_EXT = ".prox"
	VERSION  = "0.5.2"
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
