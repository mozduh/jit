/*
Copyright © 2025 mozduh
*/
package task

import (
	"github.com/spf13/cobra"
)

// taskCmd represents the task command
var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Manage Tasks and version control them",
	Long:  `Create, list, inspect, and remove tasks.`,
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
