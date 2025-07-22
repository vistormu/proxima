package evaluator

import (
	"os"
	"path/filepath"
	"strings"

	"proxima/internal/ast"
	"proxima/internal/config"

	"github.com/vistormu/go-dsa/errors"
)

type Component struct {
	name     string
	fullName string
	module   string
	content  string
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
		return nil, errors.New(MissingComponents).With(
			"components", strings.Join(missingComponents, ", "),
		)
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

func getExcludePaths(exclude []string) map[string]bool {
	excluded := map[string]bool{}
	for _, exclude := range exclude {
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
	return excluded
}

func getComponents(uniqueTags map[string]bool, config *config.Config) (map[string]Component, error) {
	// config values
	excluded := getExcludePaths(config.Components.Exclude)

	files, err := readAllFiles(config.Components.Path)
	if err != nil {
		return nil, errors.New(ReadDir).With(
			"directory", config.Components.Path,
		).Wrap(err)
	}

	// read all files
	components := map[string]Component{}
	for _, file := range files {
		file = filepath.ToSlash(filepath.Clean(file))

		if filepath.Ext(file) != ".py" {
			continue
		}

		// paths
		fullPath := filepath.Join(config.Components.Path, file)
		modules := strings.Split(strings.TrimSuffix(file, ".py"), "/")

		moduleName := ""
		if len(modules) > 1 {
			moduleName = modules[len(modules)-2]
		}

		componentName := modules[len(modules)-1]
		fullComponentName := strings.Join(modules, ".")

		// check if the file is excluded
		if _, ok := excluded[fullPath]; ok {
			continue
		}
		if _, ok := excluded[componentName]; ok {
			continue
		}
		if _, ok := excluded[fullComponentName]; ok {
			continue
		}

		// get component name
		name := componentName
		if config.Components.UseModules {
			if moduleName == componentName {
				name = strings.TrimSuffix(fullComponentName, "."+componentName)
			} else {
				name = fullComponentName
			}
		}

		// only load components that appear in the file
		if _, ok := uniqueTags[name]; !ok {
			continue
		}

		// detect if a component with the same name already exists
		if _, ok := components[name]; ok {
			return nil, errors.New(DuplicateComponent).With(
				"component", name,
				"tip", `try using the "use_modules" option in the config to avoid name conflicts`,
			)
		}

		// read file content
		content, err := getFileContent(fullPath)
		if err != nil {
			return nil, err
		}

		// check if function name and file name match
		functionName := strings.Split(content, "(")[0]
		functionName = strings.Split(functionName, "def ")[1]
		if functionName != componentName {
			return nil, errors.New(NameMismatch).With(
				"file", fullPath,
				"component", name,
				"function", functionName,
				"expected", componentName,
			)
		}

		components[name] = Component{
			name:     componentName,
			fullName: name,
			module:   moduleName,
			content:  content,
		}
	}

	return components, nil
}

func getFileContent(path string) (string, error) {
	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New(ReadFile).With(
			"file", path,
		).Wrap(err)
	}
	content := string(contentBytes)

	lines := strings.Split(content, "\n")

	var newLines []string
	var preDefLines []string

	defFound := false
	var functionIndent string

	for _, line := range lines {
		trimmedLine := strings.TrimLeft(line, " \t")
		if strings.HasPrefix(trimmedLine, "def ") && !defFound {
			defFound = true
			functionIndent = line[:len(line)-len(trimmedLine)]
			newLines = append(newLines, line)

			for _, preLine := range preDefLines {
				if preLine != "" {
					newLines = append(newLines, functionIndent+"    "+preLine)
				} else {
					newLines = append(newLines, preLine)
				}
			}
		} else {
			if !defFound {
				preDefLines = append(preDefLines, line)
			} else {
				newLines = append(newLines, line)
			}
		}
	}

	if !defFound {
		return "", errors.New(MissingDef).With(
			"file", path,
		)
	}

	return strings.Join(newLines, "\n"), nil
}
