package commands

import (
    "fmt"
    "log"
    "os"
    "time"
)

func watch(args []string) error {
    if len(args) == 0 {
        return fmt.Errorf("no files specified\nCheck 'proxima help watch' for more information")
    }
    if len(args) > 1 {
        return fmt.Errorf("watching multiple files is still not supported")
    }
    filePath := args[0]
    
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
