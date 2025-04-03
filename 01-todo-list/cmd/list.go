package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/AndrewMeleka/todo-cli/file"
	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List All Tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := listTasks()
		if err != nil {
			fmt.Println("Error listing tasks:", err)
		}
	},
}

func listTasks() error {
	// Open the file
	f, err := os.Open(file.TaskFileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	countLines, err := file.CountLines()

	if err != nil {
		fmt.Println("Error counting lines:", err)
	}

	if countLines == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	if scanner.Scan() {
	}

	// Loop through each line
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ",")

		// Format and display the line contents
		if len(fields) == 4 {
			unixTime, err := strconv.ParseInt(fields[3], 10, 64)
			if err != nil {
				fmt.Println("Error parsing timestamp:", err)
				continue
			}
			readableTime := timediff.TimeDiff(time.Unix(0, unixTime))
			fmt.Printf("ID: %s\nName: %s\nCompleted: %s\nTimestamp: %s\n\n",
				fields[0], fields[1], fields[2],
				readableTime,
			)
		} else {
			fmt.Println("Invalid data format:", line)
		}
	}

	// Check if there was an error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}

	return nil
}
