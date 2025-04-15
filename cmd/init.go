/*
Copyright © 2025 mozduh
*/
package cmd

import (
	"fmt"

	"github.com/mozduh/jit/internal/project"
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

		// check target directory for jit project
		found, path, err := project.DiscoverProject()
		if err != nil {
			fmt.Print(err)
			return
		}

		if !found {
			// initalize .jit project
			err := project.InitJitProject(path)
			if err != nil {
				fmt.Println("error init project:", err)
			}
			// print out welcom message
			fmt.Println("Created JIT Project!")
		} else {
			fmt.Println("You're already inside a jit project!")
		}

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
