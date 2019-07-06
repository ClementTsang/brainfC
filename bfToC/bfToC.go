package bftoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func lexBF(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	for {
		line, err := scanner.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatal(err)
			break
		}

		fmt.Println(fmt.Sprintf("> Read line %s", line))
		if err == io.EOF {
			break
		}
	}
}

func parseBF() {
	fmt.Println("Parsing BF...")
}

func genCCode() {

}

func writeC() {

}

// ConvertBFToC translates BF code into C code
func ConvertBFToC(inputFile string) {
	// Lex
	lexBF(inputFile)

	// Parse

	// Generate C code
}
