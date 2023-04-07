package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Dice() {
	// Set the random seed to the current time
	rand.Seed(time.Now().UnixNano())

	// Define the number of players and dice
	numPlayers := 3
	numDice := 5

	// Initialize the players and their scores
	players := make([]int, numPlayers)
	scores := make([]int, numPlayers)

	// Initialize the dice values
	dice := make([]int, numPlayers)
	fmt.Println(numDice)
	fmt.Println(players)
	fmt.Println(scores)
	fmt.Println(dice)

	// Play the game until there is only one player left
	for {
		// Each player rolls their dice
		for i := 0; i < numPlayers; i++ {
			// Check if the player has any dice left
			if players[i] < numDice {
				// Roll the player's dice
				dice[i] = rand.Intn(6) + 1

				// Check if the player rolled a 6
				if dice[i] == 6 {
					// The dice with a value of 6 is eliminated
					numDice--

					// The player gains 1 point for each 6 they rolled
					scores[i]++

					// Shift the remaining dice down to fill the gap
					copy(dice[i:], dice[i+1:numPlayers])
				}

				// Check if the player rolled a 1
				if dice[i] == 1 {
					// The next player gains an extra dice
					nextPlayer := (i + 1) % numPlayers
					numDice++
					players[nextPlayer]++
				}

				// Move to the player's next dice
				players[i]++
			}
		}

		// Check if there is only one player left with dice
		numPlayersWithDice := 0
		lastPlayer := -1
		for i, p := range players {
			if p < numDice {
				numPlayersWithDice++
				lastPlayer = i
			}
		}
		if numPlayersWithDice == 1 {
			// Declare the remaining player as the winner
			fmt.Printf("Player %d wins with a score of %d\n", lastPlayer+1, scores[lastPlayer])
			return
		}
	}
}
