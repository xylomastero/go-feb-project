package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Example finished!")
	fmt.Println("Running all Go examples")
	fmt.Println("\nAssignment 1:")
	ExampleIfForPointer()
	fmt.Println("\nAssignment 2:")
	MapsMethods()
}
