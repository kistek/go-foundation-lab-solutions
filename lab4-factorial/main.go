package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func panicHandler() {
	r := recover()

	if r == nil {
		return
	}

	fmt.Println(r)
	os.Exit(3)
}

func main() {

	const (
		inputMin = 0
		inputMax = 10
	)

	var text string
	fmt.Print("Enter factorial: ")
	fmt.Scanln(&text)

	r := regexp.MustCompile(`(?P<number>-?\d+)!`)
	match := r.FindStringSubmatch(text)

	if len(match) == 0 {
		fmt.Println("error: factorial should be in the form 'x!', e.g. '5!'")
		os.Exit(1)
	}

	number, err := strconv.Atoi(match[1])

	if err != nil {
		fmt.Println("error: converting match to number")
		os.Exit(2)
	}

	defer panicHandler()

	if number < inputMin || number > inputMax {
		// e := errors.New("error here")
		e := fmt.Errorf("error: number must be with the range [%d:%d]", inputMin, inputMax)
		panic(e)
	}

	total := 1
	for {
		if number < 2 {
			break
		}
		total = total * number
		number--
	}

	fmt.Printf("%v = %d\n", text, total)
}
