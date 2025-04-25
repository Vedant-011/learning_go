package main

import "fmt"

func main() {
	// ARRAYS AND SLICES
	fmt.Println("=== ARRAYS ===")
	// Arrays in Go have a fixed size that cannot change

	// 1. DECLARING ARRAYS
	// Format: var name [size]type
	var numbers [5]int // An array of 5 integers, all initialized to 0

	// 2. INITIALIZING ARRAYS
	var fruits [3]string = [3]string{"Apple", "Banana", "Orange"}

	// 3. SHORT DECLARATION
	colors := [4]string{"Red", "Blue", "Green", "Yellow"} // short way to make array

	// 4. LET GO COUNT THE ELEMENTS (... notation)
	animals := [...]string{"Dog", "Cat", "Bird", "Fish", "Tiger"} // Go counts 5 elements for us

	// 5. ACCESSING ELEMENTS (zero-based indexing)
	fmt.Println("First fruit:", fruits[0])  // Apple
	fmt.Println("Second color:", colors[1]) // Blue

	// 6. MODIFYING ELEMENTS
	numbers[2] = 42 // Set the third element (index 2) to 42

	// 7. ARRAY LENGTH
	fmt.Println("Number of animals:", len(animals)) // len gives how many elements

	// Print the arrays
	fmt.Println("Numbers:", numbers)
	fmt.Println("Fruits:", fruits)
	fmt.Println("Colors:", colors)
	fmt.Println("Animals:", animals)

	fmt.Println("\n=== SLICES ===")
	// Slices are like arrays but their size can change (way more flexible)

	// 1. CREATING SLICES
	var scores []int // A nil slice (no array yet) - kind of empty now
	fmt.Println("Is scores nil?", scores == nil)

	// 2. INITIALIZING SLICES
	planets := []string{"Mercury", "Venus", "Earth", "Mars"} // slice with 4 strings

	// 3. SLICES FROM ARRAYS
	someColors := colors[1:3] // Takes elenents at index 1 and 2 (Blue and Green)
	fmt.Println("Some colors:", someColors)

	// 4. MAKE FUNCTION (pre-allocate space)
	// make([]type, length, capacity)
	numbers2 := make([]int, 5, 10) // Length 5, capacity 10. will grow till 10 easily

	// 5. APPENDING ELEMENTS (growing a slice)
	scores = append(scores, 90, 85, 95)            // just adding more items like list
	planets = append(planets, "Jupiter", "Saturn") // new planets added at end

	// 6. SLICE LENGTH AND CAPACITY
	fmt.Println("Planets length:", len(planets))   // current items
	fmt.Println("Planets capacity:", cap(planets)) // total it can hold before growing again
	fmt.Println("Numbers2 length:", len(numbers2))
	fmt.Println("Numbers2 capacity:", cap(numbers2))

	// 7. MODIFYING SLICES AFFECTS UNDERLYING ARRAY
	// When you create a slice from an array, you're creating a "view" of that array
	fmt.Println("Original colors:", colors)
	someColors[0] = "NAVY" // This changes colors[1] too. (linked with  same data)
	fmt.Println("After modification:")
	fmt.Println("someColors:", someColors)
	fmt.Println("Original colors array is also affected:", colors)

	// 8. COPYING SLICES
	dest := make([]int, len(scores)) // making empty slice same size
	copied := copy(dest, scores)     // copying all items to it
	fmt.Println("Copied", copied, "elements")
	fmt.Println("Source:", scores)
	fmt.Println("Destination:", dest)

	// Print slices
	fmt.Println("Scores:", scores)
	fmt.Println("Planets:", planets)
}
