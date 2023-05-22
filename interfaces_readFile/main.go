package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, error := os.Open(os.Args[1])
	if error != nil {
		fmt.Println("Error while opening file:", error)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
