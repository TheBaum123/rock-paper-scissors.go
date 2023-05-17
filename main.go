package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var wins float64 = 0
var loses int = 0
var ties int = 0
var edgeCases int = 0
var gamesPlayed float64 = 0

func main() {
	for i := 0; i == 0; {
		var pq string
		fmt.Printf("enter p or play to play, s or stats to show stats, or anything else to quit: ")
		fmt.Scanln(&pq)
		switch pq {
		case "p", "play":
			game()
		case "s", "stats":
			logStats("max")
		default:
			fmt.Println("You decided to quit :( see you soon!")
			return
		}
	}
}

func logStats(mode string) {
	fmt.Printf("\n")
	if gamesPlayed > 0 {
		switch mode {
		case "max":
			fmt.Println("wins: ", wins)
			fmt.Println("ties: ", ties)
			fmt.Println("loses: ", loses)
			fmt.Println("total games: ", gamesPlayed)
			var winPercentage float64 = float64(math.Floor(wins/gamesPlayed*1000)) / 10
			fmt.Println("That means, you have won a total of ", winPercentage, "% of the games you played.")
		default:
			fmt.Println("You have ", wins, " wins, ", ties, " ties and ", loses, " games lost.")
			fmt.Println("You played a total of ", gamesPlayed, " games.")
			if edgeCases != 0 {
				fmt.Println("You found ", edgeCases, " edge cases. Please report them to me or open a github issue/pull request.")
			}
		}
	} else {
		fmt.Println("Please play some games first.")
	}
	fmt.Printf("\n")
}

func game() {
	var userChoice int
	var tempUserChoice string
	var computerChoice int
	var dict = [3]string{"rock", "paper", "scissors"}

	fmt.Printf("\n")

	rand.Seed(time.Now().UnixNano())
	computerChoice = rand.Intn(2) + 1

	fmt.Printf("choose rock, paper or scissors: ")
	fmt.Scanln(&tempUserChoice)

	fmt.Printf("\n")

	switch tempUserChoice {
	case "r", "rock":
		userChoice = 1
	case "p", "paper":
		userChoice = 2
	case "s", "scissors":
		userChoice = 3
	case "q", "quit":
		return
	default:
		fmt.Println("invalid choice")
		return
	}

	fmt.Println("Player chose: ", dict[userChoice-1])
	fmt.Println("Computer chose:", dict[computerChoice-1])

	fmt.Printf("\n")

	if userChoice-computerChoice == -1 || userChoice-computerChoice == 2 {
		fmt.Println("computer wins")
		loses++
	} else if userChoice-computerChoice == 0 {
		fmt.Println("its a tie")
		ties++
	} else if userChoice-computerChoice == 1 || userChoice-computerChoice == -2 {
		fmt.Println("player wins")
		wins++
	} else {
		fmt.Println("you found an edge case. please report it to me. thanks :D")
		edgeCases++
	}
	fmt.Printf("\n")
	gamesPlayed++
	logStats("default")
}
