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

func main() {
	// Demonstrate the divide function
	fmt.Println("=== RUNNING divide() ===")
	res, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	res, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", res)
	}

	// Demonstrate custom error type
	fmt.Println("\n=== CUSTOM ERROR TYPE ===")
	err = (&InvalidAgeError{Age: -3, Message: "Invalid age"}).Error()
	fmt.Println("Custom error message:", err)
}
