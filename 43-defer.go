package main

import (
	"fmt"
	"os"
)

func createFile(path string) *os.File {
	fmt.Println("creating...")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(file *os.File) {
	fmt.Println("writing...")
	fmt.Fprintln(file, "data")
}

func closeFile(file *os.File) {
	fmt.Println("closing...")
	err := file.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	f := createFile("/tmp/defer.txt")
	// defer closeFile(f)
	writeFile(f)
	closeFile(f)
}
