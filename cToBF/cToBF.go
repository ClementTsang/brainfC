package ctobf

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TODO: Add functions for interface like convertToBF()
// TODO: Move these to another file, jeez
type cToken interface {
}

func lexC(filePath string) (result []cToken) {
	return
}

func genBFCode(tokenSlice *[]*cToken) (lineSlice []*string) {
	return
}

func writeToFile(lineSlice *[]*string, filePath string) {
	dir, file := filepath.Split(filePath)
	fileToWrite := strings.TrimSuffix(file, filepath.Ext(file)) + ".c"

	// Check if it exists first, if so, delete
	if _, err := os.Stat(dir + fileToWrite); !os.IsNotExist(err) {
		os.Remove(dir + fileToWrite)
	}

	f, err := os.Create(dir + fileToWrite)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	w := bufio.NewWriter(f)

	// TODO: Always start with writing prologue:
	prologue := ""
	w.WriteString(prologue)
	w.Flush()

	for _, line := range *lineSlice {
		_, err := w.WriteString(*line)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}

	// TODO: End with writing epilogue:
	epilogue := "\n"
	w.WriteString(epilogue)
	w.Flush()
}

// ConvertCToBF translates C code into BF code
func ConvertCToBF(inputFile string, toOptimize bool, toFormat bool) {
	// Lex
	var lineArray []*string
	var tokenSlice []*cToken

	// Generate BF code and write
	lineArray = genBFCode(&tokenSlice)
	writeToFile(&lineArray, inputFile)
}
