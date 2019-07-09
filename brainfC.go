package main

import (
	bftoc "brainfC/bfToC"
	ctobf "brainfc/cToBF"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("Must include a file to convert or an argument.")
	}

	/*  TODO LIST:
	* Add optimization flags
	* Redesign flag reading system.  Currently will break on multiple flags.
	* Add input and output flags?
	* Add "I want formatting" flag
	* Add debug flag
	 */

	switch flag := os.Args[1:2][0]; flag {
	case "--help":
		fallthrough
	case "-h":
		// Give help documentation
	case "-c":
		filePath := os.Args[2:3][0]
		bftoc.ConvertBFToC(filePath, true, true)
	case "-b":
		filePath := os.Args[2:3][0]
		ctobf.ConvertCToBF(filePath, true, true)
	default:
		fmt.Println("Please input a valid command.  Type 'brainfC --help` for a list of valid flags.")
	}
}
