package cmd

import (
	"sync"

	"github.com/AndrewMeleka/todo-cli/file"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Delete All Tasks, reset the task-id counter",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			file.DeleteFile(file.TaskFileName)
		}()
		go func() {
			defer wg.Done()
			file.DeleteFile(file.TrackIDFileName)
		}()
		wg.Wait()
	},
}
