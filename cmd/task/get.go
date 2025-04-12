/*
Copyright © 2025 mozduh
*/
package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// taskGetCmd represents the taskGet command
var TaskGetCmd = &cobra.Command{
	Use:   "get <task-id>",
	Args:  cobra.ExactArgs(1),
	Short: "Inspect a task you have created",
	Long: `Inspect a task you have created and see details 
	such as the name and description of the task.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskGet called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskGetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskGetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
