package commands

import (
    "fmt"
    "os"
    "path/filepath"
    "strings"
)

func exitOnError(msg string) {
    fmt.Println("\x1b[31m-> |build| " + msg + "\x1b[0m")
    os.Exit(1)
}

func dirExists(path string) bool {
    info, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return false
        }
    }
    return info.IsDir()
}

func getAllFiles(dirPath string) []string {
    files := []string{}
    err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() && strings.HasSuffix(info.Name(), MAIN_EXT) {
            files = append(files, path)
        }
        return nil
    })
    if err != nil {
        exitOnError(err.Error())
    }
    return files
}
