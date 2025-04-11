package main

import "fmt"

// A function that takes no parameters and returns nothing
func sayHello() {
	fmt.Println("konnichiwa!")
}

// A function that takes parameters
func add(a int, b int) {
	fmt.Println("Sum:", a+b)
}

// A function that returns a value
func multiply(a int, b int) int {
	return a * b
}

// A function that returns multiple values
func divideAndRemainder(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {
	sayHello()

	add(3, 4)

	result := multiply(5, 6)
	fmt.Println("Product:", result)

	quotient, remainder := divideAndRemainder(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
}


/*
konnichiwa!
Sum: 7
Product: 30
Quotient: 3
Remainder: 1
*/

