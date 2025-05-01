package main

import (
	"fmt"
	"math/rand"
)

func guess() {
	const number = 80
	for {
		var guessed_number = rand.Intn(100) + 1
		if guessed_number > number {
			fmt.Printf(" The guessed number %v is greater than the number %v by %v \n", guessed_number, number, guessed_number-number)
		} else if number > guessed_number {
			fmt.Printf(" The guessed number %v is lesser than the number %v by %v \n", guessed_number, number, number-guessed_number)
		} else {
			fmt.Printf(" The computer guessed the number correctly which is %v ", guessed_number)
			break
		}
	}
}
