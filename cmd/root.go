/*
Copyright © 2025 mozduh
*/
package cmd

import (
	"os"

	"github.com/mozduh/jit/cmd/task"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "jit",
	Short: "CLI for tasks",
	Long:  `version control your tasks.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	importTaskCommands()
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.jit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// add commands for task toolkit
func importTaskCommands() {
	task.TaskCmd.AddCommand(task.TaskCreateCmd)
	task.TaskCmd.AddCommand(task.TaskLsCmd)
	task.TaskCmd.AddCommand(task.TaskGetCmd)
	task.TaskCmd.AddCommand(task.TaskRmCmd)

	rootCmd.AddCommand(task.TaskCmd)
}
