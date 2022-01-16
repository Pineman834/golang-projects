package main

import (
	"fmt"
	"time"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to the Regents-Physics Command-Line study helper")

	// Name and date
	fmt.Print("Whats your name => ")
	var name string
	fmt.Scan(&name)
	currentTime := time.Now()
	Date := currentTime.Format("01-02-2006")
	fmt.Printf("Todays date : %v", Date+"\n")

	// Start quiz
	fmt.Print("Ready to start the quiz y/n => ")
	var startquiz string
	fmt.Scan(&startquiz)
	if startquiz == "y" || startquiz == "Y" {
		fmt.Print("\033[H\033[2J")
	} else if startquiz == "n" || startquiz == "N" {
		return
	} else {
		fmt.Print("Answer not recognized! (Answer was not y/Y/n/N)")
		return
	}
	score := 0
	num_questions := 5

	// Question 1 : 518
	fmt.Print("A plane travelling 534 mi/hr encounters a headwind of 16 mi/hr. Determine the Resultant velocity(answer must be an integer ex: 500). \n")
	fmt.Print("=> ")
	var answer1 string
	fmt.Scan(&answer1)
	if answer1 == "518" || answer1 == "518 mi/hr" {
		fmt.Print("Correct!")
		score++
	} else {
		fmt.Print("Incorrect! The correct answer was 518 mi/hr\n")
		fmt.Print("Headwind slows down the plane which means you would subtract the magnitude of each velocity to get the resultant")
	}
	fmt.Print("\n\n")

	// Question 2 : 3
	fmt.Println("A ball is thrown horizontally with an initial velocity of 20.0 meters per second from the top of a tower 60.0 meters high.")
	fmt.Println("What is the horizontal velocity of the ball just before it reaches the ground(answer must be integer ex: 1,2,3,4)?")
	fmt.Println("1) 34.3 m/s   2) 68.6 m/s\n3) 20.0 m/s   4) 9.81 m/s")
	fmt.Print("=> ")
	var answer2 int
	fmt.Scan(&answer2)
	if answer2 == 3 {
		fmt.Print("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer was 3) 20.0 m/s")
		fmt.Println("Since we ignore air resistance in Regents-Physics there are no forces acting on the horizontal velocity meaning the velocity right before the ball hits the ground will be equal to the initial horizontal velocity.")
	}
	fmt.Print("\n\n")

	// Question 3 : 5.59 : 16
	fmt.Println("A motorboat traveling 5 m/s East encounters a current travelling 2.5 m/s North")
	fmt.Println("a. What is the resultant velocity of the motorboat? (round to nearest hundredths)")
	fmt.Print("=> ")
	var answer31 string
	fmt.Scan(&answer31)
	if answer31 == "5.59" || answer31 == "5.59 m/s" {
		fmt.Print("Correct!\n")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer was 5.59 m/s")
		fmt.Print("Use the Pythagorean Theorem(a^2+b^2=c^2) to find the resultant velocity")
	}
	fmt.Println("b. If the width of the river is 80 meters wide, then how long will it take for the boat to travel shore to shore?")
	fmt.Println("(answer should be magnitude with or without unit ex: 10, 10s)")
	fmt.Print("=> ")
	var answer32 string
	fmt.Scan(&answer32)
	if answer32 == "16" || answer32 == "16s" {
		fmt.Print("Correct!")
		score++
	} else {
		fmt.Println("Incorrect! the correct answer was 16s")
		fmt.Print("To find the time you need to divide the East velocity by 80(80/5 = 16)")
	}
	fmt.Print("\n\n")

	// Question 4 : 1
	fmt.Println("Which mass would have the greatest acceleration if the same unbalanced force was applied to each(ex: 1, A, 2, B)?")
	fmt.Println("A) 1 kg   B) 2 kg   C) 3 kg   D) 4 kg")
	fmt.Print("=> ")
	var answer4 string
	fmt.Scan(&answer4)
	if answer4 == "a" || answer4 == "A" || answer4 == "1" || answer4 == "1kg" || answer4 == "1 kg" || answer4 == "A) 1 kg" {
		fmt.Print("Correct!\n")
		score++
	} else {
		fmt.Println("Incorrect! The correct answer was A) 1 kg")
		fmt.Print("Less weight/mass will result in the most acceleration!\n")
	}

	// % correct
	percentcorrect := (float64(score) / float64(num_questions)) * 100
	fmt.Printf("You scored: %v%%\n", percentcorrect)
	fmt.Printf("You got %v out of %v\n", score, num_questions)

}
