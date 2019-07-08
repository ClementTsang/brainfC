package bftoc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// rightArrowToken represents the ">" token
type rightArrowToken struct {
	times int
}

func (r *rightArrowToken) convertToC() (result string) {
	result = fmt.Sprintf("p+=%d;\n", r.times)
	return
}
func (r *rightArrowToken) getTokenStr() (result string) {
	result = strings.Repeat(">", r.times)
	return
}
func (r *rightArrowToken) getTimes() int {
	return r.times
}
func (r *rightArrowToken) addTime() {
	r.times++
}

// leftArrowToken represents the "<" token
type leftArrowToken struct {
	times int
}

func (l *leftArrowToken) convertToC() (result string) {
	result = fmt.Sprintf("p-=%d;\n", l.times)
	return
}
func (l *leftArrowToken) getTokenStr() (result string) {
	result = strings.Repeat("<", l.times)
	return
}
func (l *leftArrowToken) getTimes() int {
	return l.times
}
func (l *leftArrowToken) addTime() {
	l.times++
}

// plusToken represents the "+" token
type plusToken struct {
	times int
}

func (p *plusToken) convertToC() (result string) {
	result = fmt.Sprintf("*p+=%d;\n", p.times)
	return
}
func (p *plusToken) getTokenStr() (result string) {
	result = strings.Repeat("+", p.times)
	return
}
func (p *plusToken) getTimes() int {
	return p.times
}
func (p *plusToken) addTime() {
	p.times++
}

// minusToken represents the "-" token
type minusToken struct {
	times int
}

func (m *minusToken) convertToC() (result string) {
	result = fmt.Sprintf("*p-=%d;\n", m.times)
	return
}
func (m *minusToken) getTokenStr() (result string) {
	result = strings.Repeat("-", m.times)
	return
}
func (m *minusToken) getTimes() int {
	return m.times
}
func (m *minusToken) addTime() {
	m.times++
}

// dotToken represents the "." token
type dotToken struct{}

func (d *dotToken) convertToC() (result string) {
	result = "putchar(*p);\n"
	return
}
func (d *dotToken) getTokenStr() (result string) {
	result = "."
	return
}

func (d *dotToken) getTimes() int {
	return 1
}

func (d *dotToken) addTime() {}

// commaToken represents the "," token
type commaToken struct{}

func (c *commaToken) convertToC() (result string) {
	result = "*p = getchar();\n"
	return
}

func (c *commaToken) getTokenStr() (result string) {
	result = ","
	return
}

func (c *commaToken) getTimes() int {
	return 1
}

func (c *commaToken) addTime() {}

// leftBracketToken represents the "[" token
type leftBracketToken struct{}

func (l *leftBracketToken) convertToC() (result string) {
	result = "while (*p) {\n"
	return
}

func (l *leftBracketToken) getTokenStr() (result string) {
	result = "["
	return
}

func (l *leftBracketToken) getTimes() int {
	return 1
}

func (l *leftBracketToken) addTime() {}

// rightBracketToken represents the "]" token
type rightBracketToken struct{}

func (r *rightBracketToken) convertToC() (result string) {
	result = "}\n"
	return
}

func (r *rightBracketToken) getTokenStr() (result string) {
	result = "]"
	return
}

func (r *rightBracketToken) getTimes() int {
	return 1
}

func (r *rightBracketToken) addTime() {}

type bfToken interface {
	convertToC() string
	getTokenStr() string
	getTimes() int
	addTime()
}

func lexBF(filePath string) (result []bfToken) {
	inCommentLoop, beforeFirstMove, currNumOfBrackets := false, true, 0 // This is required as BF code can have comments at the start that can contain command characters
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

		if rune == '[' && beforeFirstMove && !inCommentLoop {
			inCommentLoop = true
		} else if inCommentLoop {
			if rune == '[' {
				// add to num of '['s
				currNumOfBrackets++
			} else if rune == ']' {
				// subtract from num of '['s if possible, if you can't exit comment state!
				currNumOfBrackets--
				if currNumOfBrackets < 0 {
					inCommentLoop = false
				}
			}
		} else {
			beforeFirstMove = false
			switch rune {
			case '>':
				result = append(result, &rightArrowToken{1})
			case '<':
				result = append(result, &leftArrowToken{1})
			case '+':
				result = append(result, &plusToken{1})
			case '-':
				result = append(result, &minusToken{1})
			case '.':
				result = append(result, &dotToken{})
			case ',':
				result = append(result, &commaToken{})
			case '[':
				result = append(result, &leftBracketToken{})
			case ']':
				result = append(result, &rightBracketToken{})
			}
		}

		if err == io.EOF {
			break
		}
	}

	return
}

func genCCode(bfTokens *[]bfToken) (lineSlice []*string) {
	lineSlice = make([]*string, 0, 2)
	for _, token := range *bfTokens {
		line := token.convertToC()
		lineSlice = append(lineSlice, &line)
		//fmt.Printf("%s", token.convertToC())
		//fmt.Printf("%s", token.getTokenStr())
	}
	return
}

func optimizeBFToC(bfTokens *[]bfToken) {
	// Simple peephole optimization.
	// * Remove redundant increment operations (ie: +- gets reduced to nothing, >< reduced to nothing, etc)
	var prevToken *bfToken
	canOverwriteIndex := 0
	for _, token := range *bfTokens {
		(*bfTokens)[canOverwriteIndex] = token
		canOverwriteIndex++
		if prevToken != nil {
			//fmt.Printf("PREV: %c, CURR: %c\n", token.getTokenStr(), (*prevToken).getTokenStr())
			switch (*prevToken).(type) {
			case *rightArrowToken:
				switch token.(type) {
				case *leftArrowToken:
					canOverwriteIndex -= 2
				}
			case *leftArrowToken:
				switch token.(type) {
				case *rightArrowToken:
					canOverwriteIndex -= 2
				}
			case *plusToken:
				switch token.(type) {
				case *minusToken:
					canOverwriteIndex -= 2
				}
			case *minusToken:
				switch token.(type) {
				case *plusToken:
					canOverwriteIndex -= 2
				}
			}
		}
		if canOverwriteIndex-1 >= 0 {
			prevToken = &(*bfTokens)[canOverwriteIndex-1]
		} else {
			prevToken = nil
		}
	}
	// Reduce slice to whatever is left.
	*bfTokens = (*bfTokens)[:canOverwriteIndex]

	// * Combine multiple increment operations (ie: +++ normally is +p; 3 times, now make it p+=3;)
	prevToken = nil
	canOverwriteIndex = 0
	for _, token := range *bfTokens {
		(*bfTokens)[canOverwriteIndex] = token
		canOverwriteIndex++
		if prevToken != nil {
			switch (*prevToken).(type) {
			case *rightArrowToken:
				switch token.(type) {
				case *rightArrowToken:
					canOverwriteIndex--
					(*prevToken).addTime()
				}
			case *leftArrowToken:
				switch token.(type) {
				case *leftArrowToken:
					canOverwriteIndex--
					(*prevToken).addTime()
				}
			case *plusToken:
				switch token.(type) {
				case *plusToken:
					canOverwriteIndex--
					(*prevToken).addTime()
				}
			case *minusToken:
				switch token.(type) {
				case *minusToken:
					canOverwriteIndex--
					(*prevToken).addTime()
				}
			}
		}
		if canOverwriteIndex-1 >= 0 {
			prevToken = &(*bfTokens)[canOverwriteIndex-1]
		} else {
			prevToken = nil
		}
	}
	// Reduce slice to whatever is left.
	*bfTokens = (*bfTokens)[:canOverwriteIndex]

	// * Dead code removal if possible
}

func formatCCode(lineSlice *[]*string) {
	currentIndentLevel := 1
	for _, line := range *lineSlice {

		if strings.Contains(*line, "{") {
			*line = strings.Repeat("\t", currentIndentLevel) + *line
			currentIndentLevel++
		} else if strings.Contains(*line, "}") {
			currentIndentLevel--
			*line = strings.Repeat("\t", currentIndentLevel) + *line
		} else {
			*line = strings.Repeat("\t", currentIndentLevel) + *line
		}
	}
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

	// Always start with writing prologue:
	prologue := "#include <stdio.h>\nint main() {\nchar t[30000] = {0};\nchar *p = t;\n"
	w.WriteString(prologue)
	w.Flush()

	for _, line := range *lineSlice {
		_, err := w.WriteString(*line)
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
func ConvertBFToC(inputFile string, toOptimize bool, toFormat bool) {
	var lineArray []*string
	tokenSlice := lexBF(inputFile)
	// Lex
	if toOptimize {
		optimizeBFToC(&tokenSlice)
	}

	// Generate C code and write
	lineArray = genCCode(&tokenSlice)
	formatCCode(&lineArray)
	writeToFile(&lineArray, inputFile)
}
