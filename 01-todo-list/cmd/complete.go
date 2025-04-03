package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AndrewMeleka/todo-cli/file"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete task [task-id]",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error converting task ID:", err)
			return
		}

		// Get the value of the -f flag, if set
		statusFlag, _ := cmd.Flags().GetBool("false")

		// Complete the task with the appropriate status
		err = completeTask(taskId, !statusFlag) // If -f is set, make status false
		if err != nil {
			fmt.Println("Error completing task:", err)
			return
		}
	},
}

func completeTask(taskId int, status bool) (err error) {
	// Open the file for read and write (and truncate it)
	f, err := file.LoadFile(file.TaskFileName, os.O_RDWR|os.O_CREATE, 0644)
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

	// Update the task with the specified ID
	var updatedLines []string
	for _, line := range lines {
		fields := strings.Split(line, ",")
		if len(fields) > 0 {
			id, _ := strconv.Atoi(fields[0])
			if id == taskId {
				// Set the task completion status based on the flag
				fields[2] = strconv.FormatBool(status) // true or false
				// Update the timestamp
				fields[3] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
				// Join the fields back into a line
				line = strings.Join(fields, ",")
			}
			updatedLines = append(updatedLines, line)
		}
	}

	// Rewind the file and truncate it
	if err := f.Truncate(0); err != nil {
		fmt.Println("Error truncating file:", err)
		return err
	}
	if _, err := f.Seek(0, 0); err != nil {
		fmt.Println("Error seeking to the beginning of the file:", err)
		return err
	}

	// Write the updated lines back to the file
	err = file.WriteLines(f, updatedLines)
	if err != nil {
		fmt.Println("Error writing lines:", err)
		return err
	}

	return nil
}

func initCompleteCmd() {
	// Add -f flag to the command to set the status to false
	completeCmd.Flags().BoolP("false", "f", false, "Mark the task as not completed (false)")
	rootCmd.AddCommand(completeCmd)

	// Add the completeCmd to the root command or other parent commands
	// For example:
	// rootCmd.AddCommand(completeCmd)
}
