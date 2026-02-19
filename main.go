package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: hamlet64 <filename>")
		os.Exit(1)
	}

}
