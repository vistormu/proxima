package cmd

// import (
// 	"io/fs"
// 	"os"
// 	"path/filepath"
//
// 	"proxima/internal/assets"
// 	"proxima/internal/config"
//
// 	"github.com/spf13/cobra"
// 	"github.com/vistormu/go-dsa/errors"
// )
//
// var initCmd = &cobra.Command{
// 	Use:   "init",
// 	Short: "initialize a new proxima project",
// 	Long:  `the init command initializes a new proxima project by creating the necessary files and directories`,
// 	Args:  cobra.ExactArgs(0),
// 	RunE: func(cmd *cobra.Command, args []string) error {
// 		// create proxima.toml
// 		proximaTomlPath := "proxima.toml"
// 		data, err := assets.DefaultConfig.ReadFile(proximaTomlPath)
// 		if err != nil {
// 			return err
// 		}
//
// 		if err := os.MkdirAll(filepath.Dir(proximaTomlPath), 0o755); err != nil {
// 			return err
// 		}
//
// 		if err := os.WriteFile(proximaTomlPath, data, fs.FileMode(0o644)); err != nil {
// 			return err
// 		}
//
// 		// create components directory
// 		err = os.Mkdir("components", 0755)
// 		if err != nil {
// 			return errors.New(WriteFile).With(
// 				"directory", "components",
// 			).Wrap(err)
// 		}
//
// 		// create a python component
// 		f, err = os.Create("components/proxima.py")
// 		if err != nil {
// 			return errors.New(WriteFile).With(
// 				"filename", "components/proxima.py",
// 			).Wrap(err)
// 		}
// 		defer f.Close()
//
// 		_, err = f.WriteString("def proxima() -> str:\n    return \"hello from proxima!\"\n")
//
// 		// create main file
// 		f, err = os.Create("main" + MAIN_EXT)
// 		if err != nil {
// 			return errors.New(WriteFile).With(
// 				"filename", "main"+MAIN_EXT,
// 			).Wrap(err)
// 		}
// 		defer f.Close()
//
// 		_, err = f.WriteString("@proxima{}")
// 		if err != nil {
// 			return errors.New(WriteFile).With(
// 				"filename", "main"+MAIN_EXT,
// 			).Wrap(err)
// 		}
//
// 		return nil
// 	},
// }
