package examples

import "fmt"

// Animal represents basic animal properties
type Animal struct {
	Name   string
	Origin string
}

// Speak makes the animal speak
func (a *Animal) Speak() string {
	return fmt.Sprintf("%s makes a sound", a.Name)
}

// Bird embeds Animal and adds bird-specific properties
type Bird struct {
	Animal        // Embedded anonymous field
	CanFly  bool
	Speed   float64
}

// DemoEmbedding demonstrates struct embedding in Go
func DemoEmbedding() {
	fmt.Println("\n=== Struct Embedding Demo ===")

	// Create a bird
	b := Bird{
		Animal: Animal{
			Name:   "Emu",
			Origin: "Australia",
		},
		CanFly: false,
		Speed:  48, // km/h
	}

	// Access embedded fields directly
	fmt.Printf("Name: %s\n", b.Name) // Embedded field promoted
	fmt.Printf("Origin: %s\n", b.Origin)
	fmt.Printf("Can fly: %v\n", b.CanFly)
	fmt.Printf("Speed: %.1f km/h\n", b.Speed)

	// Call method from embedded type
	fmt.Println(b.Speak())

	// Shadowing example
	hawk := Bird{
		Animal: Animal{"Hawk", "North America"},
		CanFly: true,
		Speed:  390,
	}
	hawk.Name = "Red-tailed Hawk" // Modifies the embedded field
	fmt.Printf("%s can fly at %.1f km/h\n", hawk.Name, hawk.Speed)
}
