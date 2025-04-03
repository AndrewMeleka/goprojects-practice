package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"syscall"
)

const TrackIDFileName = "last_id.txt"
const TaskFileName = "tasks.csv"

func CreateFile() {
	if _, err := os.Stat(TaskFileName); !os.IsNotExist(err) {
		return
	}
	// File does not exist, create it
	file, err := os.Create(TaskFileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"ID", "Name", "Completed", "Created At"}

	if err := writer.Write(header); err != nil {
		fmt.Println("Error writing header:", err)
		return
	}
	if err := writer.Error(); err != nil {
		fmt.Println("Error flushing writer:", err)
	}
}

func CountLines() (int, error) {
	// Open the CSV file
	file, err := os.Open(TaskFileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return 0, err
	}

	// Assuming the first record is a header, so we subtract 1 from the length
	return (len(records) - 1), nil
}

func ReadLines(f *os.File) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func WriteLines(f *os.File, lines []string) error {
	// Truncate the file before writing
	if err := f.Truncate(0); err != nil {
		return err
	}
	if _, err := f.Seek(0, 0); err != nil {
		return err
	}

	for _, line := range lines {
		if _, err := f.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func LoadFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(filepath, flag, perm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	// Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func CloseFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func DeleteFile(fileName string) error {
	if err := os.Remove(fileName); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
