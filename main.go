package main

import (
	"fmt"
	"math/rand"
	"time")

func main () {
	rand.Seed(time.Now().UnixNano()) //init random numbers
	counter, guess, _number := 0, -1, rand.Intn(10)

	for counter < 3 && guess !=_number{
		fmt.Print("Guess the numbers 0..9 :")
		fmt.Scanln(&guess)
		if guess !=_number{
			counter++
			if guess < _number {
				fmt.Println("Your number is less", "")
			}else {
				fmt.Print("Your number is greater", "")
			}
		}
	}
	if guess == _number {
		fmt.Print("You WON!")
	} else {
		fmt.Print("You LOST","=", _number)
	}
}
