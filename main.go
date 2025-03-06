package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := Start(ctx); err != nil {
			log.Printf("Server error: %v", err)
		}
	}()

	fmt.Println("Running all Go examples")

	fmt.Println("\nAssignment 1:")
	ExampleIfForPointer()

	fmt.Println("\nAssignment 2:")
	MapsMethods()
	<-stop
	fmt.Println("Shutdown signal received")
	cancel()

	fmt.Println("Main function exiting")
}
