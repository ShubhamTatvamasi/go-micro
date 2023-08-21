// Generate random number
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 100
	randomNumber := rand.Intn(100)

	// Print the random number
	fmt.Println(randomNumber)
}
