package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userChoice int
	var tempUserChoice string
	var computerChoice int
	var dict = [3]string{"rock", "paper", "scissors"}

	rand.Seed(time.Now().UnixNano())
	computerChoice = rand.Intn(2) + 1

	fmt.Printf("choose rock, paper or scissors: ")
	fmt.Scanln(&tempUserChoice)

	switch tempUserChoice {
	case "r", "rock":
		userChoice = 1
	case "p", "paper":
		userChoice = 2
	case "s", "scissors":
		userChoice = 3
	case "q", "quit":
		fmt.Println("You decided to quit :( see you soon!")
		return
	default:
		fmt.Println("invalid choice")
		return
	}

	fmt.Println("Player chose: ", dict[userChoice-1])
	fmt.Println("Computer chose:", dict[computerChoice-1])

	if userChoice-computerChoice == -1 {
		fmt.Println("computer wins")
	} else if userChoice-computerChoice == 0 {
		fmt.Println("its a tie")
	} else if userChoice-computerChoice == 1 {
		fmt.Println("player wins")
	} else if userChoice-computerChoice == 2 {
		fmt.Println("computer wins")
	} else if userChoice-computerChoice == -2 {
		fmt.Println("player wins")
	} else {
		fmt.Println("you found an edge case. please report it to me. thanks :D")
	}
}
