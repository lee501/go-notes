package examples

import "fmt"

// Animal represents a basic animal
type Animal struct {
	Name string
}

// Speak makes the animal speak
func (a *Animal) Speak() {
	fmt.Printf("%s makes a sound\n", a.Name)
}

// Dog embeds Animal and adds dog-specific behavior
type Dog struct {
	Animal
	Breed string
}

// Bark is a method specific to Dog
func (d *Dog) Bark() {
	fmt.Printf("%s the %s barks!\n", d.Name, d.Breed)
}

// Override the Speak method for Dog
func (d *Dog) Speak() {
	fmt.Printf("%s the %s says: Woof!\n", d.Name, d.Breed)
}

// DemoInheritance demonstrates composition-based inheritance in Go
func DemoInheritance() {
	fmt.Println("\n=== Composition-based Inheritance ===")

	// Create a basic animal
	animal := Animal{Name: "Generic Animal"}
	animal.Speak()

	// Create a dog
	dog := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "Golden Retriever",
	}

	// Call methods
	dog.Speak()  // Calls Dog's Speak
	dog.Bark()   // Calls Dog's Bark

	// Access embedded field directly
	fmt.Printf("Dog's name: %s\n", dog.Name)

	// Call the original Speak method from Animal
	dog.Animal.Speak()

	// Interface demonstration
	var speaker interface{ Speak() } = &dog
	speaker.Speak() // Calls Dog's Speak through interface
}
