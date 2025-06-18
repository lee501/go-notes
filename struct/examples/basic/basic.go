package examples

import "fmt"

// Person represents a person with basic information
type Person struct {
	Name string
	Age  int
}

// NewPerson creates a new Person instance
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Greet makes the person introduce themselves
func (p *Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old.", p.Name, p.Age)
}

// DemoBasic demonstrates basic struct usage
func DemoBasic() {
	fmt.Println("\n=== Basic Struct Demo ===")
	
	// Value type
	var p1 Person
	p1.Name = "Alice"
	p1.Age = 25
	fmt.Printf("Value type: %+v\n", p1)


	// Pointer type with struct literal
	p2 := &Person{"Bob", 30}
	fmt.Printf("Pointer type: %+v\n", *p2)


	// Using constructor
	p3 := NewPerson("Charlie", 35)
	fmt.Printf("Using constructor: %+v\n", *p3)


	// Using new()
	p4 := new(Person)
	p4.Name = "Diana"
	p4.Age = 40
	fmt.Printf("Using new(): %+v\n", *p4)


	// Using methods
	fmt.Println(p1.Greet())
}
