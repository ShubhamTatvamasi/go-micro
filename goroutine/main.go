package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		// Perform some work
		fmt.Println("Goroutine is doing some work.")
		time.Sleep(time.Second)
		done <- true // Signal that the goroutine is done
	}()

	// Wait for the goroutine to finish by receiving from the channel
	<-done

	fmt.Println("Main function has completed.")
}
