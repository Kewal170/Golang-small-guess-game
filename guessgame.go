package main

import (
	f "fmt"
	"math/rand"
	t "time"
)

func main() {
	var userInput int
	rand.Seed(t.Now().UnixNano())

	winNum := rand.Intn(10)

	f.Println("Welcome to Guess the Game with Golang!")
	f.Println("Enter your guess :> ")
	//f.Printf("Answer is %v\n", winNum)
	f.Scanln(&userInput)

	if userInput > 10 {
		f.Printf("Please enter below 10")
	} else {
		if userInput == winNum {
			f.Println("You won")
		} else {
			f.Println("You lose")
		}
	}
}
