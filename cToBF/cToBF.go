package ctobf

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// TODO: Add functions for interface like convertToBF()

type cToken struct {
	cTokenType string
}

func genBFCode(tokenSlice []*cToken) (lineSlice []*string) {
	return
}

func getCToken(stringGrouping string) (result *cToken) {

	return
}

func isCSymbol(currentRune rune) (result bool) {

	const listOfCSymbols string = "=+-*/&^%|.?!;{}[]<>():"

	result = strings.Contains(listOfCSymbols, string(currentRune))
	return
}

func isValidMultiCharSymbol(currentStringGrouping string, currentRune rune, nextRune rune) (result bool) {
	switch currentStringGrouping {
	case ".":
		result = (currentRune == '.' && nextRune == '.')
	case ">>":
		fallthrough
	case "<<":
		fallthrough
	case "=":
		fallthrough
	case "/":
		fallthrough
	case "*":
		fallthrough
	case "^":
		fallthrough
	case "%":
		fallthrough
	case "!":
		result = (currentRune == '=')
	case "+":
		result = (currentRune == '+' || currentRune == '=')
	case "-":
		result = (currentRune == '-' || currentRune == '=' || currentRune == '>')
	case "|":
		result = (currentRune == '|' || currentRune == '=')
	case "&":
		result = (currentRune == '&' || currentRune == '=')
	case "<":
		result = (currentRune == '<' || currentRune == '=' || currentRune == ':' || currentRune == '%')
	case ">":
		result = (currentRune == '>' || currentRune == '=' || currentRune == ':' || currentRune == '%')
	default:
		result = false
	}

	//fmt.Printf("IsValidMultiCharSymbol with: %s, %c, %c - result: %t\n", currentStringGrouping, currentRune, nextRune, result)

	return
}

func getTokensForLine(line string) (result []*cToken) {

	// Format and remove unneeded parts:
	line = regexp.MustCompile("//.*").ReplaceAllString(line, "")
	fmt.Printf("Regexed string: %s\n", line)

	// Crawl by character!
	line += " " // Cheat by appending a newline, so the last character in the line will also be output
	currentStringGrouping := ""
	for pos, currentChar := range line {
		if unicode.IsSpace(currentChar) {

			if currentStringGrouping != "" {
				fmt.Printf("Grouping: %s\n", currentStringGrouping)
			}

			currentStringGrouping = ""
		} else if isCSymbol(currentChar) {
			if currentStringGrouping == "" || (currentStringGrouping != "" && pos < len(line) && isValidMultiCharSymbol(currentStringGrouping, currentChar, []rune(line)[pos+1])) {
				// Check to see if this character can continue to form a valid token with the currentStringGrouping
				currentStringGrouping += string(currentChar)

				if !isCSymbol([]rune(line)[pos+1]) {
					fmt.Printf("Grouping: %s\n", currentStringGrouping)
					currentStringGrouping = ""
				}

			} else {
				// Parse currentStringGrouping and current value of char
				if currentStringGrouping != "" {
					fmt.Printf("Grouping: %s\n", currentStringGrouping)
				}

				fmt.Printf("Char: %s\n", string(currentChar))

				currentStringGrouping = ""
			}
		} else {
			currentStringGrouping += string(currentChar)
		}
	}

	return
}

func lexC(filePath string) (result []*cToken) {
	result = make([]*cToken, 0, 2)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	for {
		str, err := scanner.ReadString('\n') // read every line
		if err != nil && err != io.EOF {
			log.Fatal(err)
			break
		}

		str = strings.TrimSpace(str) // trim newlines

		// Read every string as a token here
		result = append(result, getTokensForLine(str)...)

		if err == io.EOF {
			break
		}
	}

	return
}

func writeToFile(lineSlice *[]*string, filePath string) {
	dir, file := filepath.Split(filePath)
	fileToWrite := strings.TrimSuffix(file, filepath.Ext(file)) + ".bf"

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

	tokenSlice = lexC(inputFile)

	// Generate BF code and write
	lineArray = genBFCode(tokenSlice)
	writeToFile(&lineArray, inputFile)
}
