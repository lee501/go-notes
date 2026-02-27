package examples

import (
	"encoding/json"
	"fmt"
	"log"
)

// User demonstrates struct tags for JSON marshaling
type User struct {
	ID        int    `json:"id,omitempty"`
	Username  string `json:"username"`
	Email     string `json:"email,omitempty"`
	Active    bool   `json:"active"`
	CreatedAt string `json:"created_at,omitempty"`
	password  string // Unexported fields are not marshaled
}

// DemoJSON demonstrates JSON marshaling and unmarshaling with structs
func DemoJSON() {
	fmt.Println("\n=== JSON Marshaling Demo ===")

	// Create a user
	user := User{
		ID:        1,
		Username:  "johndoe",
		Email:     "john@example.com",
		Active:    true,
		password:  "secret123",
		CreatedAt: "2023-01-01T00:00:00Z",
	}

	// Marshal to JSON
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling to JSON: %v", err)
	}

	fmt.Println("Marshaled JSON:")
	fmt.Println(string(jsonData))

	// Unmarshal from JSON
	var newUser User
	jsonStr := `{"id":2,"username":"janedoe","email":"jane@example.com","active":true}`

	err = json.Unmarshal([]byte(jsonStr), &newUser)
	if err != nil {
		log.Fatalf("Error unmarshaling from JSON: %v", err)
	}

	fmt.Println("\nUnmarshaled struct:")
	fmt.Printf("%+v\n", newUser)
}
