/*
Copyright © 2025 mozduh
*/
package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// taskCreateCmd represents the taskCreate command
var TaskCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create a task!",
	Long:  `create a task with a name and a description`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskCreate called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
