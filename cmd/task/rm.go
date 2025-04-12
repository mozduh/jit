/*
Copyright © 2025 mozduh
*/
package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// taskRmCmd represents the taskRm command
var TaskRmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove a task.",
	Long:  `Remove a task you have created.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskRm called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskRmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskRmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
