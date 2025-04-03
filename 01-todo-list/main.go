package main

import (
	"github.com/AndrewMeleka/todo-cli/cmd"
	"github.com/AndrewMeleka/todo-cli/file"
)

func main() {
	file.CreateFile()
	// create a file if it doesn't exist
	cmd.Execute()
}
