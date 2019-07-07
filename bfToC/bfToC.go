package bftoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

func lexBF(filePath string) []bfToken {
	result := make([]bfToken, 0, 2)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	for {
		nonWhiteSpace := true
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
		default:
			nonWhiteSpace = false
		}

		if nonWhiteSpace {
			fmt.Printf("Saw a %c token.\n", result[len(result)-1].getTokenChar())
		}

		if err == io.EOF {
			break
		}
	}

	return result
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
