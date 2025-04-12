/*
Copyright © 2025 mozduh
*/

package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

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
