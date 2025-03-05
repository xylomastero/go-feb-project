package main

import "fmt"

func main() {
	defer fmt.Println("Example finished!")
	fmt.Println("Running all Go examples")

	// Start the API in a separate goroutine
	go func() {
		fmt.Println("Starting API server...")
		Start()
	}()

	// Run other examples
	fmt.Println("\nAssignment 1:")
	ExampleIfForPointer()
	fmt.Println("\nAssignment 2:")
	MapsMethods()

	// Keep the main goroutine running
	select {}
}
