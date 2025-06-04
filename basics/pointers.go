package main

import "fmt"

// Function that takes a pointer to an int
func doubleValue(x *int) {
	// Dereference the pointer to modify the actual value
	*x = *x * 2
}

// Function that doesn't use pointers (just for comparison)
func doubleValueNonPointer(x int) int {
	return x * 2
}

// Struct for pointer demonstration
type Counter struct {
	value int
}

// Method that uses a pointer receiver
func (c *Counter) increment() {
	c.value++ // Automatically dereferenced
}

// Method that doesn't use a pointer receiver (just for comparison)
func (c Counter) incrementNonPointer() Counter {
	c.value++
	return c
}

func main() {
	fmt.Println("=== POINTERS IN GO ===")
	// A pointer holds the memory address of a value

	// 1. CREATING POINTERS
	x := 42
	// The & operator creates a pointer to a variable
	var p *int = &x // p is a pointer to the int x

	fmt.Println("Value of x:", x)
	fmt.Println("Memory address of x (value of p):", p)

	// 2. DEREFERENCING POINTERS
	// The * operator accesses the value a pointer points to
	fmt.Println("Value that p points to:", *p) // This is called "dereferencing"

	// 3. MODIFYING VALUES THROUGH POINTERS
	*p = 100 // Change the value of x through the pointer p
	fmt.Println("New value of x after modifying through pointer:", x)

	// 4. ZERO VALUE OF POINTERS
	var nilPointer *int // Zero value of a pointer is nil
	fmt.Println("Nil pointer:", nilPointer)

	// Uncommenting the line below would cause a crash..
	// fmt.Println(*nilPointer) // Error: panic: runtime error: invalid memory address or nil pointer dereference

	// 5. POINTER TO POINTER
	var pp **int = &p // pp points to p which points to x..
	fmt.Println("Pointer to pointer value:", pp)
	fmt.Println("Value obtained through double dereferencing:", **pp) // 100

	// 6. FUNCTION USING POINTERS
	y := 5
	fmt.Println("Original y:", y)
	doubleValue(&y)                                      // Pass the address of y
	fmt.Println("y after doubleValue function call:", y) // 10

	// Compareing with non-pointer version
	z := 5
	fmt.Println("Original z:", z)
	result := doubleValueNonPointer(z)                                 // Pass a copy of z
	fmt.Println("z after doubleValueNonPointer function call:", z)     // Still 5
	fmt.Println("Result from doubleValueNonPointer function:", result) // 10

	// 7. POINTERS WITH STRUCTS
	counter := Counter{value: 0}
	fmt.Println("Initial counter value:", counter.value)

	// Using pointer receiver method
	ptrCounter := &counter // Create a pointer to counter
	ptrCounter.increment()
	fmt.Println("Counter value after increment with pointer:", counter.value) // 1

	// Go allows automatic dereferencing for struct pointers
	ptrCounter.increment()                                               // Same as (*ptrCounter).increment()
	fmt.Println("Counter value after another increment:", counter.value) // 2

	// Compare with non-pointer receiver method
	newCounter := counter.incrementNonPointer()                     // This doesn't modify the original counter
	fmt.Println("Original counter value:", counter.value)           // Still 2
	fmt.Println("New counter value from return:", newCounter.value) // 3

	// 8. NEW FUNCTION
	// The new function creates a pointer to a zero-valued instance
	ptr := new(int) // Create a pointer to a zero-valued int (0)
	fmt.Println("Value at ptr:", *ptr)
	*ptr = 42
	fmt.Println("New value at ptr:", *ptr)

	// WHY USE POINTERS?
	fmt.Println("\nWHY USE POINTERS?")
	fmt.Println("1. To modify variables across function calls")
	fmt.Println("2. To avoid copying large data structures (efficiency)")
	fmt.Println("3. To express that a variable might be optional (nil)")
	fmt.Println("4. To implement certain data structures like linked lists")
}
