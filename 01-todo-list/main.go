package main

import (
	"github.com/AndrewMeleka/goprojects-practice/cmd"
	"github.com/AndrewMeleka/goprojects-practice/file"
)

func main() {
	file.CreateFile()
	// create a file if it doesn't exist
	cmd.Execute()
}
