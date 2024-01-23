package main

import (
	"fmt"
	"os"
)

func getFileContents(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("test")
	// Put code here
}
