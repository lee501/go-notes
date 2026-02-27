package Struct

import "fmt"

// People demonstrates basic method definition and embedding
type People struct{}

// ShowA demonstrates method calling another method on the same receiver
func (p *People) ShowA() {
	fmt.Println("People's ShowA")
	p.ShowB()
}

// ShowB is a method that can be overridden by embedded types
func (p *People) ShowB() {
	fmt.Println("People's ShowB")
}

// Teachers embeds People to demonstrate method overriding
type Teachers struct {
	People
}

// ShowB overrides the ShowB method from the embedded People type
func (t *Teachers) ShowB() {
	fmt.Println("Teachers' ShowB")
}

// DemoInheritance demonstrates composition-based inheritance in Go
func DemoInheritance() {
	t := &Teachers{}
	fmt.Println("\n=== Composition-based Inheritance ===")
	fmt.Println("Calling t.ShowA():")
	t.ShowA() // Calls People's ShowA which calls the overridden ShowB

	fmt.Println("\nCalling t.ShowB():")
	t.ShowB() // Directly calls Teacher's ShowB
}
