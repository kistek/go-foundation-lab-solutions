package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	const maxGuesses = 5

	unixTime := time.Now().UnixNano()
	seed := rand.NewSource(unixTime)
	prg := rand.New(seed)
	correctInt := prg.Intn(100)

	guesses := []int{}
	attempt := 0

	for {
		var guess string
		fmt.Print("Enter guess: ")
		fmt.Scanln(&guess)

		guessInt, err := strconv.Atoi(guess)
		if err != nil {
			fmt.Println("Could not convert guess to integer")
			continue
		}

		guesses = append(guesses, guessInt)

		if guessInt == correctInt {
			fmt.Println("Got it!")
			break
		}

		attempt++

		if attempt >= maxGuesses {
			fmt.Println("No more guesses, game over.")
			fmt.Printf("The correct number was %d\n", correctInt)
			break
		}

		if guessInt > correctInt {
			fmt.Println("too high")
		} else {
			fmt.Println("too low")
		}

	}

	fmt.Printf("Here were your guesses: %v\n", guesses)
}
