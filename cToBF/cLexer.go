package ctobf

import (
	"bufio"
	"io"
	"log"
	"os"
)

type cToken interface {
	getCombinedRegexValue() string
}

type runeComponents interface {
	getRegexValue() string
}

func getToken(str string) (result cToken) {
	return
}

func lexC(filePath string) (result []cToken) {
	result = make([]cToken, 0, 2)
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewReader(file)
	for {
		str, err := scanner.ReadString(' ')
		if err != nil && err != io.EOF {
			log.Fatal(err)
			break
		}

		result = append(result, getToken(str))

		if err == io.EOF {
			break
		}
	}

	return
}
