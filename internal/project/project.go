/*
Copyright © 2025 mozduh
*/
package project

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mozduh/jit/internal/task"
)

const (
	ProjectDirName  = ".jit"
	ProjectFileName = "config.json"
)

var defaultProjectConfig = map[string]interface{}{
	"version": "0.1.0",
	"counter": 0,
}

func InitJitProject(cwd string) error {

	projectDirPath := filepath.Join(cwd, ProjectDirName)
	projectFilePath := filepath.Join(projectDirPath, ProjectFileName)
	taskDirBacklog := filepath.Join(cwd, task.TaskDirBacklog)
	taskDirTodo := filepath.Join(cwd, task.TaskDirTodo)
	taskDirProgress := filepath.Join(cwd, task.TaskDirProgress)
	taskDirDone := filepath.Join(cwd, task.TaskDirDone)

	// create .jit directory
	if err := os.MkdirAll(projectDirPath, 0755); err != nil {
		return err
	}

	// create task bin directories
	if err := os.MkdirAll(taskDirBacklog, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(taskDirTodo, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(taskDirProgress, 0755); err != nil {
		return err
	}
	if err := os.MkdirAll(taskDirDone, 0755); err != nil {
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

	if err := os.WriteFile(projectFilePath, data, 0644); err != nil {
		return err
	}

	return nil
}
