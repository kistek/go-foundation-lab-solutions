package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatal("usage: fib start_int end_int")
	}

	start, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal("Could not read start integer arg!")
	}

	end, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Could not read end integer arg")
	}

	if end < 1 || end > 100 {
		end = 100
		fmt.Println("maximum end value is 100, using that value instead")
	}

	if start < 1 || start > 100 {
		log.Fatal("start must be between 1 and 100")
	}

	var fibs = []int{0, 1}

	for c := 0; c < end; c++ {
		nextVal := fibs[c] + fibs[c+1]
		fibs = append(fibs, nextVal)
	}

	fmt.Println(fibs[start-1 : end])
}
