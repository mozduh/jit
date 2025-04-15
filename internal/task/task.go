/*
Copyright © 2025 mozduh
*/
package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Task struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Column      string `json:"column"`
}

const TaskDirBacklog = "backlog"
const TaskDirTodo = "todo"
const TaskDirProgress = "progress"
const TaskDirDone = "done"

func (t *Task) Save(savePath string) error {

	path := filepath.Join(savePath, fmt.Sprintf("%s.json", t.ID))
	data, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadTask(dir, id string) (*Task, error) {
	path := filepath.Join(dir, fmt.Sprintf("%s.json", id))
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var t Task
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func LoadAllTasks(dir string) ([]Task, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			data, err := os.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				return nil, err
			}
			var t Task
			if err := json.Unmarshal(data, &t); err != nil {
				return nil, err
			}
			tasks = append(tasks, t)
		}
	}

	return tasks, nil
}

func (t *Task) Delete(dir string) error {
	path := filepath.Join(dir, fmt.Sprintf("%s.json", t.ID))
	return os.Remove(path)
}
