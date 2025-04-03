package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/AndrewMeleka/todo-cli/file"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := createTask(args[0])
		if err != nil {
			fmt.Println("Error creating task:", err)
		}
	},
}

type Task struct {
	ID        int
	Name      string
	Completed bool
	Timestamp int64
}

func createTask(name string) (err error) {
	f, err := file.LoadFile(file.TaskFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	if err != nil {
		fmt.Println("Error counting lines:", err)
		return err
	}

	// Create a new task with the current ID, name, and timestamp
	id, err := getNextID()
	if err != nil {
		fmt.Println("Error getting next ID:", err)
	}
	now := time.Now().UTC().UnixNano()
	task := Task{
		ID:        id,
		Name:      name,
		Completed: false,
		Timestamp: now,
	}

	// Convert task struct to CSV format
	record := []string{
		strconv.Itoa(task.ID),
		task.Name,
		strconv.FormatBool(task.Completed),
		strconv.FormatInt(task.Timestamp, 10),
	}

	if err := writer.Write(record); err != nil {
		fmt.Println("Error writing task:", err)
		return err
	}

	if err := saveLastID(task.ID); err != nil {
		fmt.Println("Error saving last ID:", err)
	}

	fmt.Println("Task added successfully!")
	return nil
}

func getNextID() (int, error) {
	idFile := "last_id.txt"

	// Try to read the last ID
	data, err := os.ReadFile(idFile)
	if err != nil {
		// If the file doesn't exist, start from 1
		if os.IsNotExist(err) {
			return 1, nil
		}
		return 0, err
	}

	lastID, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}

	return lastID + 1, nil
}

func saveLastID(id int) error {
	return os.WriteFile("last_id.txt", []byte(strconv.Itoa(id)), 0644)
}
