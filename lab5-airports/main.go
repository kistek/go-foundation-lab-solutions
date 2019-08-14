package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var text string
	fmt.Print("Enter route: ")

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("failed to read from stdin")
		os.Exit(1)
	}

	trimmed := strings.TrimSpace(text)
	codes := strings.Split(trimmed, " ")

	var names []string

	for _, code := range codes {

		name, exists := IataToName[code]

		if !exists {
			fmt.Printf("Could not find airport with code: %q\n", code)
			os.Exit(2)
		}

		names = append(names, name)
	}

	fmt.Println(strings.Join(names, " to "))

}
