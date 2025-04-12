/*
Copyright © 2025 mozduh
*/
package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

// taskLsCmd represents the taskLs command
var TaskLsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List tasks you have created",
	Long: `List tasks you have created and see a table of 
	id field along with name and description.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("taskLs called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// taskLsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// taskLsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
