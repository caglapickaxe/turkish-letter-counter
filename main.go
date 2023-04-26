package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"unicode"
)

func createMap(chars string) map[rune]struct{} {
	res := make(map[rune]struct{}, len(chars))
	for _, r := range chars {
		res[r] = struct{}{}
	}
	return res
}

func mapRune(r rune) rune {
	switch r {
	case 'â', 'Â':
		return 'A'
	case 'î', 'Î':
		return 'İ'
	case 'û', 'Û':
		return 'U'
	default:
		return unicode.TurkishCase.ToUpper(r)
	}
}

var runes = createMap("AÂBCÇDEFGĞHIİÎJKLMNOÖPRSŞTUÛÜVYZ")

var inputName = flag.String("input", "", "input file")
var outputName = flag.String("output", "out.txt", "output file")

func main() {
	flag.Parse()

	inputFile, err := os.Open(*inputName)
	if err != nil {
		log.Fatalln("failed to open input file:", err)
	}

	freqs := make(map[rune]int, len(runes))
	input := bufio.NewReader(inputFile)
	for {
		if r, _, err := input.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln("failed to read input:", err)
			}
		} else {
			mapped := mapRune(r)
			if _, ok := runes[mapped]; ok {
				freqs[mapped]++
			}
		}
	}
	inputFile.Close()

	outputFile, err := os.OpenFile(*outputName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatalln("failed to open output file:", err)
	}
	output := bufio.NewWriter(outputFile)

	for r, freq := range freqs {
		output.WriteRune(r)
		output.WriteString(": ")
		output.WriteString(strconv.Itoa(freq))
		output.WriteByte('\n')
	}

	output.Flush()
	outputFile.Close()
}
