package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AndrewMeleka/todo-cli/file"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete Task [task-id]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid task ID:", err)
			return
		}

		err = deleteTask(taskId)
		if err != nil {
			fmt.Println("Error deleting tasks:", err)
		}
	},
}

func deleteTask(taskId int) error {
	// Open the file
	f, err := os.OpenFile(file.TaskFileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()

	// Read all lines into a slice
	lines, err := file.ReadLines(f)
	if err != nil {
		fmt.Println("Error reading lines:", err)
		return err
	}

	// Remove the task with the specified ID
	var updatedLines []string
	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) > 0 {
			id, _ := strconv.Atoi(fields[0])
			if id != taskId {
				updatedLines = append(updatedLines, line)
			}
		}
	}

	// Write the updated lines back to the file
	err = file.WriteLines(f, updatedLines)
	if err != nil {
		fmt.Println("Error writing lines:", err)
		return err
	}

	return nil

}
