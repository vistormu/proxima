package commands

import (
    "fmt"
    "os"
    "time"
    "strings"
)

func watch(args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("no files specified\nCheck 'proxima help watch' for more information")
    }
    if len(args) > 1 {
        return fmt.Errorf("watching multiple files is still not supported")
    }
    filePath := args[0]

    if !strings.HasSuffix(filePath, MAIN_EXT) {
        return fmt.Errorf("file %s is not a .prox file", filePath)
    }
    
    ticker := time.NewTicker(2 * time.Second)
    defer ticker.Stop()

    fileInfo, err := os.Stat(filePath)
    if err != nil {
        return err
    }
    lastModTime := fileInfo.ModTime()

    msg := fmt.Sprintf("\x1b[32m-> |watch| Watching file %s for changes...\x1b[0m", filePath)
    fmt.Println(msg)

    for range ticker.C {
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            return err
        }
        currentModTime := fileInfo.ModTime()
        if currentModTime.After(lastModTime) {
            err := generate([]string{filePath})
            if err != nil {
                fmt.Println("\x1b[31m-> |build| Error transpiling file:\x1b[0m\n", err)
            }

            lastModTime = currentModTime
        }
    }

    return nil
}
