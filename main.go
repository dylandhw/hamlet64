package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var totalPadding = 0


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
	bitString := bytesToBitString(byteSlice)

	paddedBitString := padBitString(bitString)

}

func bytesToBitString(byteSlice []byte) string {
	var builder strings.Builder

	// iterate through every byte in the byte slice
	for _, b := range byteSlice { 
		builder.WriteString(fmt.Sprintf("%08b", b))
	}

	return builder.String()
}

func padBitString(bitString string) string {
	paddingNeeded := (6 - (len(bitString) % 6)) % 6
	totalPadding += paddedNeeded 

	bitString.Repeat("0", paddingNeeded)

	return bitString
}