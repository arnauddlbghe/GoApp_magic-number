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

	var userInput int
	for {
		fmt.Print("Try a number : ")
		fmt.Scan(&userInput)
		fmt.Printf("Testing %d\n", userInput)

		if userInput == magicNumber {
			fmt.Println("Victory")
			return
		}
		if userInput > magicNumber {
			fmt.Println("Your number is too big")
			continue
		}
		if userInput < magicNumber {
			fmt.Println("Your number is too small")
			continue
		}
	}
}
