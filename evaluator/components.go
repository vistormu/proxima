package evaluator

import (
    "os"
    "strings"
    "path/filepath"

    "proxima/config"
    "proxima/ast"
    "proxima/errors"
)

type Component struct {
    name string
    fullName string
    content string
}

func loadComponents(expressions []ast.Expression, config *config.Config) (map[string]Component, error) {
    // get all unique tags
    uniqueTags := map[string]bool{}
    getUniqueTags(expressions, uniqueTags)

    // get all components
    components, err := getComponents(uniqueTags, config)
    if err != nil {
        return nil, err
    }

    // check that all tags have a corresponding component
    missingComponents := []string{}
    for tag := range uniqueTags {
        if _, ok := components[tag]; !ok {
            missingComponents = append(missingComponents, tag)
        }
    }
    if len(missingComponents) > 0 {
        return nil, errors.NewComponentError(errors.MISSING_COMPONENTS, missingComponents)
    }

    return components, nil
}

func getUniqueTags(expressions []ast.Expression, uniqueTags map[string]bool) {
    for _, expression := range expressions {
        switch expression := expression.(type) {
        case *ast.Tag:
            if _, ok := uniqueTags[expression.Name]; !ok {
                uniqueTags[expression.Name] = true
            }
            for _, arg := range expression.Args {
                getUniqueTags(arg.Values, uniqueTags)
            }
        }
    }
}

func readAllFiles(dir string) ([]string, error) {
    return readAllFilesLoop(dir, dir)
}

func readAllFilesLoop(baseDir, dir string) ([]string, error) {
    entries, err := os.ReadDir(dir)
    if err != nil {
        return nil, err
    }

    var files []string
    for _, entry := range entries {
        fullPath := filepath.Join(dir, entry.Name())
        if entry.IsDir() {
            subFiles, err := readAllFilesLoop(baseDir, fullPath)
            if err != nil {
                return nil, err
            }
            files = append(files, subFiles...)
        } else {
            relativePath, err := filepath.Rel(baseDir, fullPath)
            if err != nil {
                return nil, err
            }
            files = append(files, relativePath)
        }
    }

    return files, nil
}

func isDir(path string) bool {
    fileInfo, err := os.Stat(path)
    if err != nil {
        return false
    }
    return fileInfo.IsDir()
}

func getComponents(uniqueTags map[string]bool, config *config.Config) (map[string]Component, error) {
    // config values
    componentsDir := config.Components.Path
    useModules := config.Components.UseModules
    excluded := map[string]bool{}
    for _, exclude := range config.Components.Exclude {
        if isDir(exclude) {
            files, err := readAllFiles(exclude)
            if err != nil {
                continue
            }
            for _, file := range files {
                filename := filepath.Join(exclude, file)
                excluded[filename] = true
            }
            continue
        }
        excluded[exclude] = true
    }

    // get all files under the components directory recursively
    files, err := readAllFiles(componentsDir)
    if err != nil {
        return nil, errors.NewComponentError(errors.ERROR_READING_DIR, componentsDir, err.Error())
    }

    // read all files
    components := map[string]Component{}
    for _, file := range files {
        // paths
        fullPath := filepath.Join(componentsDir, file)
        componentName := strings.Split(filepath.Base(file), ".")[0]
        fullComponentName := strings.ReplaceAll(strings.Split(file, ".")[0], "/", ".")

        // check if the file is excluded
        if _, ok := excluded[fullPath]; ok {
            continue
        }

        // only load python files
        if filepath.Ext(fullPath) != ".py" {
            continue
        }
        
        // get component name
        name := componentName
        if useModules {
            name = fullComponentName
        }

        // only load components that appear in the file
        if _, ok := uniqueTags[name]; !ok {
            continue
        }

        // detect if a component with the same name already exists
        if _, ok := components[name]; ok {
            return nil, errors.NewComponentError(errors.DUPLICATE_COMPONENT, name)
        }

        // read file content
        content, err := getFileContent(fullPath)
        if err != nil {
            return nil, err
        }

        components[name] = Component{
            name: componentName,
            fullName: name,
            content: content,
        }
    }

    return components, nil
}

func getFileContent(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return "", errors.NewComponentError(errors.ERROR_READING_FILE, path, err.Error())
    }
    return string(content), nil
}
