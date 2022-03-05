package main

import (
	"fmt"
	"math/rand"
)

func main() {
	turns := 10
	fmt.Println("How many turns would you like?")
	fmt.Print("| => ")
	fmt.Scan(&turns)
	game(turns)
}

func game(turns int) {

	score := 0
	for i := 0; i < turns; i++ {

		num1 := rand.Intn(10)
		num2 := rand.Intn(10)

		fmt.Println("What is", num1, "+", num2, "?")
		fmt.Print("Your answer: ")
		var answer int
		fmt.Scan(&answer)

		if answer == num1+num2 {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Incorrect!")
			if score == 0 {
			} else {
				score--
			}
		}
	}

	fmt.Println("\nYour score is", score, "out of", turns)
}
