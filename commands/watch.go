package commands

import (
    "fmt"
    "log"
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
    
    ticker := time.NewTicker(2 * time.Second) // Check every 2 seconds
    defer ticker.Stop()

    fileInfo, err := os.Stat(filePath)
    if err != nil {
        log.Fatal(err)
    }
    lastModTime := fileInfo.ModTime()

    for range ticker.C {
        fileInfo, err := os.Stat(filePath)
        if err != nil {
            log.Println("Error accessing file:", err)
            continue
        }
        currentModTime := fileInfo.ModTime()
        if currentModTime.After(lastModTime) {
            log.Println("File modified, transpiling...")
            err := generate([]string{filePath})
            if err != nil {
                log.Println("Error transpiling file:", err)
            }

            lastModTime = currentModTime // Update the last modified time
        }
    }

    return nil
}
