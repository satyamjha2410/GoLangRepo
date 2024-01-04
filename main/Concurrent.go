package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	// Start a Goroutine that executes the printNumbers function concurrently
	go printNumbers()

	// Let the main Goroutine do its work
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Main Goroutine is working...")
	}

	// Wait for a key press to allow the Goroutine to finish (for demonstration purposes)
	fmt.Println("\nPress Enter to exit.")
	fmt.Scanln()
	fmt.Println("Exiting...")
}
