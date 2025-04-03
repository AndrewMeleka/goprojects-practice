package cmd

import (
	"fmt"
	"os"

	"github.com/AndrewMeleka/goprojects-practice/file"
	"github.com/spf13/cobra"
)

var deleteAllCmd = &cobra.Command{
	Use:   "delete-all",
	Short: "Delete All Tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		err := deleteTasks()
		if err != nil {
			println("Error deleting tasks:", err)
		}
	},
}

func deleteTasks() error {
	// Open the file
	f, err := os.OpenFile(file.FileName, os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()

	// // Read all lines into a slice
	lines, err := file.ReadLines(f)
	if err != nil {
		fmt.Println("Error reading lines:", err)
		return err
	}

	// Write the updated lines back to the file
	err = file.WriteLines(f, lines[:1])
	if err != nil {
		fmt.Println("Error writing lines:")
		return err
	}

	return nil

}
