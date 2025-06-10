package main

import (
	"errors"
	"fmt"
)

// Function that can return an error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// Create and return a new error
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil // Return result and nil error
}

// Custom error type
type InvalidAgeError struct {
	Age     int
	Message string
}

// Implement the Error interface for our custom error type
func (e *InvalidAgeError) Error() string {
	return fmt.Sprintf("%s: %d is not a valid age", e.Message, e.Age)
}

// Function using our custom error type
func validateAge(age int) error {
	if age < 0 {
		return &InvalidAgeError{
			Age:     age,
			Message: "Negative age provided",
		}
	}
	if age > 150 {
		return &InvalidAgeError{
			Age:     age,
			Message: "Age too high",
		}
	}
	return nil // no error here
}

func main() {
	fmt.Println("=== Running Snippet 1 + 2 ===")

	// Test divide
	result, err := divide(100, 5)
	if err != nil {
		fmt.Println("Divide error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	// Test validateAge
	err = validateAge(160)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age is valid")
	}
}
