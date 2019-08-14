package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	oneLine := strings.ReplaceAll(fileString, "\n", " ")
	rawWords := strings.Split(oneLine, " ")

	var wordCount = map[string]int{}

	for _, r := range rawWords {
		str := string(r)
		str = strings.TrimSpace(str)
		if str != "" {
			wordCount[str]++
		}
	}

	for word, count := range wordCount {
		fmt.Printf("%s %d\n", word, count)
	}
}
