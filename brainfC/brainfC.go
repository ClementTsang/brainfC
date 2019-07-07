package main

import (
	bftoc "brainfC/bfToC"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Must include a file to convert or an argument.")
	}

	switch flag := os.Args[1:2][0]; flag {
	case "--help":
		fallthrough
	case "-h":
		// Give help documentation
	case "-c":
		filePath := os.Args[2:3][0]
		bftoc.ConvertBFToC(filePath)
	case "-b":
	default:
	}
}
