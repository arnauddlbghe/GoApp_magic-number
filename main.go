package main

import (
	"fmt"
	"math/rand/v2"
)

// 1. Generate RandomNumber
// 2. Loop asking for user to insert a number (Validate that it is a number)
// 3. Correct: Stop
// 4. Not correct:
// 4.1 : User gives a bigger number: Notify him
// 4.2 : User gives a smaller number: Notify him.
func main() {
	var magicNumber = rand.IntN(100)
	fmt.Println("MagicNumber generated. Good game !")

	remainingTries := 6
	for remainingTries > 0 {
		fmt.Printf("%d attempts lefts. Try a number : ", remainingTries)
		userInput := getUserInput()
		fmt.Printf("Testing %d\n", userInput)

		if userInput == magicNumber {
			fmt.Println("Victory")
			return
		} else if userInput > magicNumber {
			fmt.Println("Your number is too big")
		} else if userInput < magicNumber {
			fmt.Println("Your number is too small")
		}

		remainingTries--
		if remainingTries == 0 {
			fmt.Printf("You lost the game, it was %d. Wanna play again ?\n", magicNumber)
		}
	}
}

func getUserInput() int {
	var userInput int
	for {
		fmt.Scan(&userInput)
		if userInput > 0 {
			return userInput
		}
		fmt.Println("Invalid number. Number need to be positive (> 0)")
	}
}
