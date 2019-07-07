package main

import (
	bftoc "brainfC/bfToC"
	"os"
)

func main() {
	filePath := os.Args[1:2]
	bftoc.ConvertBFToC(filePath[0])
}
