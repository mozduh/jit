/*
Copyright © 2025 mozduh
*/
package project

import (
	"errors"
	"os"
	"path/filepath"
)

// DiscoverRepo walks up from the current directory to find a `.mydir` folder.
// It returns the absolute path to the directory containing `.mydir`.
func DiscoverProject() (bool, string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return false, "", err
	}

	for {
		dotDir := filepath.Join(dir, ".jit")
		info, err := os.Stat(dotDir)
		if err == nil && info.IsDir() {
			return true, dir, nil
		}

		// Stop at root
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return false, "", errors.New(".jit directory not found in current or parent directories")
}
