package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filenames := os.Args[1:]

	if len(filenames) == 0 {
		fmt.Println("No files supplied")
		os.Exit(1)
	}

	for i, filename := range filenames {
		file, err := os.Open(filename)

		if err != nil {
			fmt.Println("Error with opening file", filename)
			os.Exit(1)
		}

		if i != 0 {
			fmt.Println("")
		}

		io.Copy(os.Stdout, file)

	}
	// filename := filenames[0]

}
