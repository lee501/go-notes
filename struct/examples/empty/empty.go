package examples

import (
	"fmt"
	"sync"
)

// EmptySet demonstrates using a map with empty struct values as a set
type EmptySet map[string]struct{}


// DemoEmptyStruct demonstrates the use of empty structs
func DemoEmptyStruct() {
	fmt.Println("\n=== Empty Struct Demo ===")

	// 1. Channel signaling
	done := make(chan struct{})

	// Start a goroutine that will signal when done
	go func() {
		// Do some work...
		fmt.Println("  Goroutine: Work done")
		close(done) // Signal completion by closing the channel
	}()

	// Wait for the signal
	_, ok := <-done
	if !ok {
		fmt.Println("  Main: Received done signal")
	}

	// 2. Using empty struct as a value in a map (set implementation)
	set := make(EmptySet)
	set["apple"] = struct{}{}
	set["banana"] = struct{}{}

	// Check if a value exists in the set
	if _, exists := set["apple"]; exists {
		fmt.Println("  'apple' is in the set")
	}

	// 3. Using empty struct as a method receiver
	var empty emptyType
	empty.doSomething()

	// 4. Using empty struct in a WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("  Background task completed")
	}()
	wg.Wait()
}

// emptyType demonstrates using empty struct as a method receiver
type emptyType struct{}

// doSomething is a method on an empty struct
func (e emptyType) doSomething() {
	fmt.Println("  Method called on empty struct")
}
