/*
Copyright © 2025 mozduh
*/

package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const (
	ProjectDirName  = ".jit"
	ProjectFileName = "config.json"
)

var defaultProjectConfig = map[string]interface{}{
	"version": "0.1.0",
	"counter": 0,
}

func CheckForJitDir(cwd string) bool {
	fmt.Println("Checking for .jit")

	found := false
	dir := cwd

	for {
		// Check if `.dir` exists directly inside this path
		checkPath := filepath.Join(dir, ".jit")
		info, err := os.Stat(checkPath)
		if err == nil && info.IsDir() {
			// fmt.Println("Found .jit at:", checkPath)
			found = true
			break
		}

		// Move up
		parent := filepath.Dir(dir)
		if parent == dir {
			break // Reached root
		}
		dir = parent
	}

	return found
}

func InitJitProject(cwd string) error {

	projectDirPath := filepath.Join(cwd, ProjectDirName)
	projectFilePath := filepath.Join(projectDirPath, ProjectFileName)

	// Create the .dir directory if it doesn't exist
	if err := os.MkdirAll(projectDirPath, 0755); err != nil {
		return err
	}

	// Check if config.yaml exists
	if _, err := os.Stat(projectFilePath); err == nil {
		return fmt.Errorf("config file already exists at %s", projectFilePath)
	}

	// Write default config file
	data, err := json.MarshalIndent(defaultProjectConfig, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(projectFilePath, data, 0644)
}
