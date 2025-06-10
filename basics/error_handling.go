package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
	fmt.Println("=== ERROR HANDLING IN GO ===")
	// Go uses explicit error handling with return values
	// (no exceptions like in other languges)

	// 1. BASIC ERROR HANDLING PATTERN
	result, err := divide(10, 2)
	if err != nil {
		// Handle the error
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Trying with an error case
	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error occurred:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// 2. CREATING ERRORS
	// Using errors.New
	err1 := errors.New("this is a simple error")
	fmt.Println("Simple error:", err1)

	// Using fmt.Errorf (allows formatting)
	name := "John"
	age := 30
	err2 := fmt.Errorf("%s is %d years old", name, age)
	fmt.Println("Formatted error:", err2)

	// 3. ERROR HANDLING WITH MULTIPLE FUNCTIONS
	// Convert string to number (which can fail)
	numStr := "42"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		// If successful, use the result in another function that can fail
		result, err := divide(float64(num), 2)
		if err != nil {
			fmt.Println("Error in division:", err)
		} else {
			fmt.Println("Final result:", result)
		}
	}

	// Try with an invalid number
	invalidStr := "not a number"
	_, err = strconv.Atoi(invalidStr)
	if err != nil {
		fmt.Println("Error converting invalid string:", err)
	}

	// 4. CUSTOM ERROR TYPES
	err = validateAge(25)
	if err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Age is valid")
	}

	err = validateAge(-5)
	if err != nil {
		fmt.Println("Validation error:", err)

		// Type assertion to access custom error fields
		if ageErr, ok := err.(*InvalidAgeError); ok {
			fmt.Println("This was an age error with age:", ageErr.Age)
		}
	}

	// 5. FILE OPERATIONS (COMMON ERROR HANDLING SCENARIO)
	file, err := os.Open("nonexistent-file.txt")
	if err != nil {
		fmt.Println("File error:", err)

		// Using the errors package to check specific errors
		if os.IsNotExist(err) {
			fmt.Println("The file doesn't exist - creating it")
			// Here you would create the file (not implemented in this example)
		}
	} else {
		defer file.Close() // Ensure the file is closed when we're done
		fmt.Println("File opened successfully")
		// Read from file, etc.
	}

	// 6. DEFER, PANIC, AND RECOVER
	fmt.Println("\n=== DEFER, PANIC, AND RECOVER ===")

	// DEFER: schedules a function call to run when the surrounding function returns
	defer fmt.Println("This runs after main() completes (defer)")

	// Multiple defers run in LIFO order (Last In, First Out)
	defer fmt.Println("This runs second (defer)")
	defer fmt.Println("This runs first (defer)")

	// PANIC and RECOVER
	fmt.Println("Demonstrating panic and recover:")

	// This function will demonstrate recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Comment/uncomment to see panic in action
	// panic("deliberate panic")

	fmt.Println("This code runs normally if no panic occurs")
}
