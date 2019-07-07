package bftoc

import (
	"bufio"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// rightArrowToken represents the ">" token
type rightArrowToken struct {
}

func (r rightArrowToken) convertToC() (result string) {
	result = "++ptr;\n"
	return
}

func (r rightArrowToken) getTokenChar() (result rune) {
	result = '>'
	return
}

// leftArrowToken represents the "<" token
type leftArrowToken struct {
}

func (l leftArrowToken) convertToC() (result string) {
	result = "--ptr;\n"
	return
}

func (l leftArrowToken) getTokenChar() (result rune) {
	result = '<'
	return
}

// plusToken represents the "+" token
type plusToken struct {
}

func (p plusToken) convertToC() (result string) {
	result = "++*ptr;\n"
	return
}

func (p plusToken) getTokenChar() (result rune) {
	result = '+'
	return
}

// minusToken represents the "-" token
type minusToken struct {
}

func (m minusToken) convertToC() (result string) {
	result = "--*ptr;\n"
	return
}

func (m minusToken) getTokenChar() (result rune) {
	result = '-'
	return
}

// dotToken represents the "." token
type dotToken struct {
}

func (d dotToken) convertToC() (result string) {
	result = "putchar(*ptr);\n"
	return
}

func (d dotToken) getTokenChar() (result rune) {
	result = '.'
	return
}

// commaToken represents the "," token
type commaToken struct {
}

func (c commaToken) convertToC() (result string) {
	result = "*ptr = getChar();\n"
	return
}

func (c commaToken) getTokenChar() (result rune) {
	result = ','
	return
}

// leftBracketToken represents the "[" token
type leftBracketToken struct {
}

func (l leftBracketToken) convertToC() (result string) {
	result = "while (*ptr) {\n"
	return
}

func (l leftBracketToken) getTokenChar() (result rune) {
	result = '['
	return
}

// rightBracketToken represents the "]" token
type rightBracketToken struct {
}

func (r rightBracketToken) convertToC() (result string) {
	result = "}\n"
	return
}

func (r rightBracketToken) getTokenChar() (result rune) {
	result = ']'
	return
}

type bfToken interface {
	convertToC() string
	getTokenChar() rune
}

func lexBF(filePath string) (result []bfToken) {
	result = make([]bfToken, 0, 2)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	for {
		rune, _, err := scanner.ReadRune()
		if err != nil && err != io.EOF {
			log.Fatal(err)
			break
		}

		switch rune {
		case '>':
			result = append(result, rightArrowToken{})
		case '<':
			result = append(result, leftArrowToken{})
		case '+':
			result = append(result, plusToken{})
		case '-':
			result = append(result, minusToken{})
		case '.':
			result = append(result, dotToken{})
		case ',':
			result = append(result, commaToken{})
		case '[':
			result = append(result, leftBracketToken{})
		case ']':
			result = append(result, rightBracketToken{})
		}

		if err == io.EOF {
			break
		}
	}

	return
}

func genCCode(bfTokens *[]bfToken) (lineSlice []string) {
	lineSlice = make([]string, 0, 2)
	for _, token := range *bfTokens {
		lineSlice = append(lineSlice, token.convertToC())
	}
	return
}

func optimizeBFToC() (optimizedLineSlice []string) {
	return
}

func writeToFile(lineSlice *[]string, filePath string) {
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

	// Always start with writing prologue:
	prologue := "#include <stdio.h>\nint main() {\nchar array[5000000] = {0};\nchar *ptr = array;\n"
	w.WriteString(prologue)
	w.Flush()

	for _, line := range *lineSlice {
		_, err := w.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}

	// End with writing epilogue:
	epilogue := "}\n"
	w.WriteString(epilogue)
	w.Flush()
}

// ConvertBFToC translates BF code into C code
func ConvertBFToC(inputFile string) {
	// Lex
	tokenSlice := lexBF(inputFile)

	// Generate C code and write
	lineArray := genCCode(&tokenSlice)
	writeToFile(&lineArray, inputFile)
}
