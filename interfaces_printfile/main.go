package main

import (
	"fmt"
	"io"
	"os"
)

// implements a program that reads content from a text and prints its content to the terminal.
// Use _go run main.go myfile.txt_
// Use os.Args which is a slice of string
// os.Open can help to open a file
// What interfaces does the "File" type implement?
// If the "File" type implement "Reader" interface, you might be able to reuse the io.Copy function.

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Wrong arguments, use 'go run main.go myfile.txt'")
		os.Exit(1)
	}

	filename := args[1]

	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("Something gets wrong when try to open the file:", err)
		os.Exit(2)
	}

	fmt.Println("Read file", filename)
	io.Copy(os.Stdout, file)
	fmt.Println("\n--------------------------")

}
