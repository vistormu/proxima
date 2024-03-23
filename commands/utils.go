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

func dirExists(path string) (bool, error) {
    info, err := os.Stat(path)
    if err != nil {
        if os.IsNotExist(err) {
            return false, nil
        }
        return false, err
    }
    return info.IsDir(), nil
}

func isDir(path string) (bool, error) {
    info, err := os.Stat(path)
    if err != nil {
        return false, err
    }
    return info.IsDir(), nil
}

func getAllFiles(dirPath string) ([]string, error) {
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
        return nil, err
    }
    return files, nil
}

func getFilesInDir(dir string) ([]string, error) {
    var filesWithExt []string
    entries, err := os.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }

        fileInfo, err := entry.Info()
        if err != nil {
            return nil, err
        }
        if strings.HasSuffix(fileInfo.Name(), MAIN_EXT) {
            filesWithExt = append(filesWithExt, filepath.Join(dir, fileInfo.Name()))
        }
    }

    return filesWithExt, nil
}
