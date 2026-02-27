package examples

import "fmt"

// Counter demonstrates methods with value and pointer receivers
type Counter struct {
	value int
}

// Increment demonstrates a pointer receiver method
func (c *Counter) Increment() {
	c.value++
}

// Value demonstrates a value receiver method
func (c Counter) Value() int {
	return c.value
}

// Reset resets the counter to zero
func (c *Counter) Reset() {
	c.value = 0
}

// DemoMethods demonstrates method receivers in Go
func DemoMethods() {
	fmt.Println("\n=== Methods Demo ===")

	// Create a counter
	counter := Counter{value: 0}

	// Call methods
	fmt.Printf("Initial value: %d\n", counter.Value())
	
	counter.Increment()
	counter.Increment()
	fmt.Printf("After incrementing twice: %d\n", counter.Value())

	counter.Reset()
	fmt.Printf("After reset: %d\n", counter.Value())

	// Method values and expressions
	increment := counter.Increment // Method value
	increment()                    // Equivalent to counter.Increment()
	fmt.Printf("After method value call: %d\n", counter.Value())


	// Method expression
	reset := (*Counter).Reset
	reset(&counter) // Need to pass the receiver
	fmt.Printf("After method expression call: %d\n", counter.Value())
}
