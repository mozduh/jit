/*
Copyright © 2025 mozduh
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init jit project",
	Long: `Init a .jit project directory in the current working directory.
	This directory will be used to manage the task data of this project.
	`,
	Run: func(cmd *cobra.Command, args []string) {

		// get current working directory
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("error getting current directory:", err)
		}

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

		if !found {
			// initalize .jit project
			fmt.Println("No .jit directory found in any ancestor of the current directory.")
			// INIT(cwd)

			// print out welcom message
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
