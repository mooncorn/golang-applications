package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	play(90)
}

func play(max int) {
	fmt.Printf("Welcome to Guess A Number! The generated number is between 0 and %d!\n", max)

	randNumber := rand.Intn(max + 1)
	guess := ""
	guessed := false

	for !guessed {
		fmt.Print("Take a guess: ")
		fmt.Scanln(&guess)
		num, err := strconv.Atoi(guess)

		if err != nil {
			fmt.Println("Invalid input. Try entering an integer.")
			continue
		}

		if num == randNumber {
			fmt.Println("You guessed the number!")
			guessed = true
		} else if num > randNumber {
			fmt.Println("Your guess is too high!")
		} else {
			fmt.Println("Your guess is too low!")
		}
	}
}
