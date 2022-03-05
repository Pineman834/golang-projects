package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num := rand.Intn(100)
	toolow := [4]string{"try guessing a little higher",
		"your a little too low", "Gotta guess higher",
		"think higher"}
	toohigh := [4]string{"try guessing a little higher",
		"your guess was a little too high", "Gotta guess lower",
		"think lower"}
	correct := [4]string{"you got it", "Guessed correct", "Yup thats correct", "Right on"}
	var guessedcorrect bool

	for i := 0; i < 5; i++ {
		var guess int
		fmt.Println("Guess a number 1-100")
		fmt.Print("| => ")
		fmt.Scan(&guess)

		if guess < num {
			fmt.Println(toolow[rand.Intn(3)])
		}
		if guess > num {
			fmt.Println(toohigh[rand.Intn(3)])
		}
		if guess == num {
			fmt.Println(correct[rand.Intn(3)])
			guessedcorrect = true
			break
		}
	}
	if !guessedcorrect {
		fmt.Println("your out of guesses, your number was", num)
	}

}
