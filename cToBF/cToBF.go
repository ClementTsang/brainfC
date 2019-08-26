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

// Globals

// Based on http://www.quut.com/c/ANSI-C-grammar-l-2011.html
const o = "[0-7]"
const d = "[0-9]"
const nz = "[1-9]"
const l = "[a-zA-Z_]"
const a = "[a-zA-Z_0-9]"
const h = "[a-fA-F0-9]"
const hp = "[a-fA-F0-9]"
const e = "(0[xX])"
const p = "([Pp][+-]?{D}+)"
const fs = "(f|F|l|L)"
const is = "(((u|U)(l|L|ll|LL)?)|((l|L|ll|LL)(u|U)?))"
const cp = "(u|U|L)"
const sp = "(u8|u|U|L)"
const es = "(\\\\(['\"\\?\\\\abfnrtv]|[0-7]{1,3}|x[a-fA-F0-9]+))"
const ws = "[ \t\v\n\f]"

var validKeywords = []string{
	"auto", "break", "case", "char", "const", "continue", "default", "do", "double", "else", "enum",
	"extern", "float", "for", "goto", "if", "inline", "int", "long", "register", "restrict", "return",
	"short", "signed", "sizeof", "static", "struct", "switch", "typedef", "union", "unsigned", "void",
	"volatile", "while",
}

var underscoredKeywords = []string{
	"_Alignas", "_Alignof", "_Atomic", "_Bool", "_Complex", "_Generic", "_Imaginary", "_Noreturn", "_Static_Assert", "_Thread_local",
}

const funcName = "__func__"

var validMultiCharSymbols = []string{
	"...", ">>=", "<<=", "+=", "-=", "*=", "/=", "%=", "&=", "^=", "|=", ">>", "<<", "++", "--", "->", "&&", "||", "<=", ">=", "==", "!=",
}

var validSingleCharSymbols = []string{
	";", "{", "}", "<%", "%>", ",", ":", "=", "(", ")", "]", "<:", ":>", "[", ".", "&", "!", "~", "-", "+", "*", "/", "%", "<", ">", "^", "|", "?",
}

var checkTypeRegex = regexp.MustCompile(l + (a) + "*")

var iConstantRegexOne = regexp.MustCompile(hp + h + "+" + is + "?")
var iConstantRegexTwo = regexp.MustCompile(nz + d + "*" + is + "?")
var iConstantRegexThree = regexp.MustCompile("0" + o + "*" + is + "?")
var iConstantRegexFour = regexp.MustCompile(cp + "?" + "'([^'\\\n]|" + es + ")" + "'")

var fConstantRegexOne = regexp.MustCompile(d + "+" + e + fs + "?")
var fConstantRegexTwo = regexp.MustCompile(d + "*" + "\".\"" + d + "+" + e + "?" + fs + "?")
var fConstantRegexThree = regexp.MustCompile(d + "+" + "\".\"" + e + "?" + fs + "?")
var fConstantRegexFour = regexp.MustCompile(hp + h + "+" + p + fs + "?")
var fConstantRegexFive = regexp.MustCompile(hp + h + "*" + "\".\"" + h + "+" + p + fs + "?")
var fConstantRegexSix = regexp.MustCompile(hp + h + "+" + "\".\"" + p + fs + "?")

var stringLiteralRegex = regexp.MustCompile("(" + sp + "?" + "\\\"([^\"\\\\\n]|" + es + ")*\\\"" + ws + "*)+")

//===========================================//

// TODO: Add functions for interface like convertToBF()

type cToken struct {
	actualVal  string
	cTokenType string
}

func genBFCode(tokenSlice []*cToken) (lineSlice []*string) {
	return
}

func arrayContains(arr []string, targ string) bool {
	for _, ele := range arr {
		if ele == targ {
			return true
		}
	}
	return false
}

func getCToken(stringGrouping string) (result *cToken) {
	//fmt.Printf("Looking for CToken for %s\n", stringGrouping)

	tokenType := "UNKNOWN"
	if arrayContains(validKeywords, stringGrouping) {
		tokenType = strings.ToUpper(stringGrouping)
	} else if arrayContains(underscoredKeywords, stringGrouping) {
		tokenType = strings.ToUpper(stringGrouping[1:])
	} else if stringGrouping == funcName {
		tokenType = "FUNC_NAME"
	} else if arrayContains(validMultiCharSymbols, stringGrouping) {
		tokenType = "MUTLICHAR" //TODO: Return correct multicharsymbol
	} else if arrayContains(validSingleCharSymbols, stringGrouping) {
		switch stringGrouping {
		case "<%":
			tokenType = "{"
		case "%>":
			tokenType = "}"
		case "<:":
			tokenType = "["
		case ":>":
			tokenType = "]"
		default:
			tokenType = stringGrouping
		}
	} else if checkTypeRegex.MatchString(stringGrouping) {
		tokenType = "CHKTYPE" //TODO: Checktype
	} else if iConstantRegexOne.MatchString(stringGrouping) ||
		iConstantRegexTwo.MatchString(stringGrouping) ||
		iConstantRegexThree.MatchString(stringGrouping) ||
		iConstantRegexFour.MatchString(stringGrouping) {
		tokenType = "I_CONSTANT"
	} else if fConstantRegexOne.MatchString(stringGrouping) ||
		fConstantRegexTwo.MatchString(stringGrouping) ||
		fConstantRegexThree.MatchString(stringGrouping) ||
		fConstantRegexFour.MatchString(stringGrouping) ||
		fConstantRegexFive.MatchString(stringGrouping) {
		tokenType = "F_CONSTANT"
	} else if stringLiteralRegex.MatchString(stringGrouping) {
		tokenType = "STRING_LITERAL"
	}

	result = &cToken{actualVal: stringGrouping, cTokenType: tokenType}

	return
}

func isCSymbol(currentRune rune) (result bool) {

	const listOfCSymbols string = "=+-*/&^%|.?!;{}[]<>():"

	result = strings.Contains(listOfCSymbols, string(currentRune))
	return
}

func isValidMultiCharSymbol(currentStringGrouping string, runeSlice ...rune) (result bool) {
	currentRune := runeSlice[0]
	switch currentStringGrouping {
	case ".":
		result = (currentRune == '.' && len(runeSlice) > 1 && runeSlice[1] == '.')
	case ">>":
		fallthrough
	case "<<":
		fallthrough
	case "=":
		fallthrough
	case "^":
		fallthrough
	case "%":
		fallthrough
	case "!":
		result = (currentRune == '=')
	case "*":
		result = (currentRune == '=' || currentRune == '/')
	case "/":
		result = (currentRune == '=' || currentRune == '*')
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

	//fmt.Printf("IsValidMultiCharSymbol with: %s, %v- result: %t\n", currentStringGrouping, runeSlice, result)

	return
}

func getTokensForLine(line string) (result []*cToken) {

	// Crawl by character!
	line += " " // Cheat by appending a newline, so the last character in the line will also be output
	currentStringGrouping := ""
	for pos, currentChar := range line {
		if unicode.IsSpace(currentChar) {

			if currentStringGrouping != "" {
				result = append(result, getCToken(currentStringGrouping))
			}

			currentStringGrouping = ""
		} else if isCSymbol(currentChar) {
			//fmt.Printf("Current symbol grouping: %s, current char: %s\n", currentStringGrouping, string(currentChar))

			if currentStringGrouping == "" || (currentStringGrouping != "" && pos+1 < len(line) && isValidMultiCharSymbol(currentStringGrouping, currentChar, []rune(line)[pos+1])) {
				// Check to see if this character can continue to form a valid token with the currentStringGrouping
				currentStringGrouping += string(currentChar)

				if !isCSymbol([]rune(line)[pos+1]) {
					result = append(result, getCToken(currentStringGrouping))
					currentStringGrouping = ""
				}
			} else {
				result = append(result, getCToken(currentStringGrouping))
				currentStringGrouping = string(currentChar)

				if !isCSymbol([]rune(line)[pos+1]) {
					result = append(result, getCToken(currentStringGrouping))
					currentStringGrouping = ""
				}
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

		// Instead, read as a full string, and trim all comments first!

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

	for _, token := range tokenSlice {
		fmt.Printf("Token: %s => %s\n", token.actualVal, token.cTokenType)
	}

	// Generate BF code and write
	lineArray = genBFCode(tokenSlice)
	writeToFile(&lineArray, inputFile)
}
