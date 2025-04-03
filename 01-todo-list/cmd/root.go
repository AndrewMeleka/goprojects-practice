package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cli-task",
	Short: "A simple CLI Task Management BY AN",
	Long:  "A Task Management tool to manage your tasks efficiently.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the CLI Task tool! Use --help to see available commands.")
	},
}

func Execute() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(resetCmd)
	initCompleteCmd()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
