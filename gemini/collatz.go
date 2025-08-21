// Package main provides a command-line interface to test the Collatz conjecture
// for a range of numbers using multi-threading (goroutines).
package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

// CollatzResult holds the result of a Collatz sequence calculation for a single number.
type CollatzResult struct {
	Number int64
	Steps  int
}

// collatzSequence calculates the Collatz sequence for a given positive integer n.
// It returns the number of steps to reach 1. This version is slightly simplified
// to avoid creating and returning a large slice, which is more efficient for
// a multi-threaded application.
func collatzSequence(n int64) int {
	// Initialize the steps counter.
	steps := 0
	current := n

	// Loop until the number becomes 1.
	for current != 1 {
		// Check if the number is even or odd.
		if current%2 == 0 {
			// If even, divide by 2.
			current = current / 2
		} else {
			// If odd, multiply by 3 and add 1.
			current = 3*current + 1
		}
		// Increment the steps.
		steps++
	}

	// Return the final step count.
	return steps
}

// worker is a goroutine function that calculates the Collatz steps for a number
// and sends the result to a channel.
// The wait group is used to signal when this worker is done.
func worker(id int, numbers <-chan int64, results chan<- CollatzResult, wg *sync.WaitGroup) {
	// Signal to the wait group that this worker is done when the function exits.
	defer wg.Done()
	// Loop over the incoming numbers channel until it is closed.
	for n := range numbers {
		// Calculate the steps.
		steps := collatzSequence(n)
		// Send the result back on the results channel.
		results <- CollatzResult{
			Number: n,
			Steps:  steps,
		}
	}
}

// main is the entry point of the program.
// It sets up the multi-threaded processing of a range of numbers.
func main() {
	// Check if the correct number of arguments (start and end range) were provided.
	if len(os.Args) < 3 {
		fmt.Println("Please provide a start and end positive integer to test.")
		fmt.Println("Usage: go run main.go <start> <end>")
		os.Exit(1)
	}

	// Parse the start number.
	start, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil || start <= 0 {
		fmt.Println("Error: Invalid start number. Please enter a valid positive integer.")
		os.Exit(1)
	}

	// Parse the end number.
	end, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil || end <= 0 || end < start {
		fmt.Println("Error: Invalid end number. Please enter a valid positive integer greater than or equal to the start number.")
		os.Exit(1)
	}

	// Define the number of worker goroutines to use. Using a reasonable number
	// like 4 or 8 is often a good starting point for modern CPUs.
	const numWorkers = 8

	// Create a wait group to wait for all workers to finish.
	var wg sync.WaitGroup
	// Create channels to pass numbers to workers and receive results.
	numbers := make(chan int64, numWorkers)
	results := make(chan CollatzResult, (end-start)+1)

	// Launch worker goroutines.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, numbers, results, &wg)
	}

	// Send numbers to the numbers channel.
	for n := start; n <= end; n++ {
		numbers <- n
	}
	// Close the numbers channel to signal workers that no more numbers will be sent.
	close(numbers)

	// Wait for all the workers to complete their tasks.
	wg.Wait()
	// Close the results channel after all workers are done.
	close(results)

	// Process and print the results from the results channel.
	fmt.Printf("Processing Collatz conjecture for numbers from %d to %d...\n", start, end)
	for result := range results {
		fmt.Printf("Number %d took %d steps.\n", result.Number, result.Steps)
	}
	fmt.Println("All calculations are complete.")
}
