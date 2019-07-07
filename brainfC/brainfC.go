package main

import (
	bftoc "brainfC/bfToC"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Must include a file to convert.")
	}
	filePath := os.Args[1:2]
	bftoc.ConvertBFToC(filePath[0])
}
