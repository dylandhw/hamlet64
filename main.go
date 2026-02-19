package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: hamlet64 <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Println("file: ", filename)
}

func run(filename string) {}
