package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: hamlet64 <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Println("file: ", filename)

	process(filename)
}

func process(filename string) {
	// convert file into byte slice
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("trouble reading file: %v", err)
	}
	fmt.Println("byteslice:", byteSlice)

	encode(byteSlice)
}

func encode(byteSlice []byte) string {

}
