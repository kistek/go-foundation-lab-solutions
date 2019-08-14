// An alternate solution to the lab if you want to count words only (not punctuation)
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {

	var filename string

	if len(os.Args) < 2 {
		fmt.Printf("usage: %s filename.txt\n", os.Args[0])
		return
	}

	filename = os.Args[1]

	fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileString := string(fileBytes)

	var sb strings.Builder
	var wordCount = map[string]int{}

	for _, r := range fileString {

		if unicode.IsLetter(r) {
			sb.WriteRune(r)
			continue
		}
		if sb.Len() == 0 {
			continue
		}
		word := sb.String()
		wordCount[word]++
		sb.Reset()
	}

	for word, count := range wordCount {
		fmt.Printf("%s %d\n", word, count)
	}
}
