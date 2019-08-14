package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
)

// State represents the name and size of US states
type State struct {
	Name string
	Size int
}

var states []State
var average int

// DisplayState pretty prints information about a State
func (s State) DisplayState() string {
	if s.Size > average {
		return fmt.Sprintf("%s is greater than average (%d)", s.Name, s.Size)
	}
	return fmt.Sprintf("%s is less than average (%d)", s.Name, s.Size)
}

func main() {
	fileBytes, err := ioutil.ReadFile("states.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileString := string(fileBytes)
	countriesData := strings.Split(fileString, "\n")

	var total int

	for _, value := range countriesData {

		index := strings.LastIndex(value, " ")
		state := value[0:index]
		size := value[index+1:]
		isize, _ := strconv.Atoi(strings.TrimSpace(size))
		states = append(states, State{state, isize})

		total += isize
	}

	average = total / len(states)

	fmt.Printf("Total size %d\n", total)
	fmt.Printf("Average size %d\n", average)

	t := template.New("states")
	_, err = t.Parse("{{range .	}}{{- .DisplayState}}\n{{ end }}")
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(os.Stdout, states)
	if err != nil {
		log.Fatal(err)
	}

}
