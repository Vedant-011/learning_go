package main

import "fmt"

func main() {
	// --- IF, ELSE IF, ELSE ---
	var age int = 22

	//checking age and trying to group people based on that
	if age < 13 {
		fmt.Println("You're a kid.")
	} else if age < 20 {
		fmt.Println("You're a teenager.")
	} else if age < 60 {
		fmt.Println("You're an adult.")
	} else {
		fmt.Println("You're a senior citizen.")
	}

	// --- SWITCH STATEMENT ---
	day := 3

	// trying switch here to map number to days.. 3 should be Wednesday
	switch day {
	case 1:
		fmt.Println("It's Monday")
	case 2:
		fmt.Println("It's Tuesday")
	case 3:
		fmt.Println("It's Wednesday")
		fallthrough //fallthrough makes it also do next case
	case 4:
		fmt.Println("Or maybe Thursday?")
	default:
		fmt.Println("It's some other day")
	}

	// --- FOR LOOP (classic) ---
	fmt.Println("Counting from 1 to 5:")
	for i := 1; i <= 5; i++ {
		fmt.Println(i) // simple loop
	}

	// --- FOR LOOP (while-style) ---
	count := 3
	fmt.Print("Countdown: ")
	for count > 0 {
		fmt.Print(count, " ") // just printing count till it hits 0
		count--               // decreasing it every time
	}
	fmt.Println("Lift off!")

	// --- FOR RANGE LOOP (slices/arrays) ---
	numbers := []int{10, 20, 30, 40}
	fmt.Println("Numbers and their indexes:")
	for index, value := range numbers {
		// range gives both index and value here
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// --- FOR WITH BREAK/CONTINUE ---
	fmt.Println("Odd numbers under 10:")
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue // skip even ones
		}
		fmt.Println(i)
		if i == 7 {
			break // done once hit 7
		}
	}

	// --- INFINITE LOOP (with break) ---
	fmt.Println("Simulating input until a condition is met (x == 3):")
	x := 0
	for {
		x++

		if x == 3 {
			fmt.Println("Found it! x =", x)
			break // stopping it
		}
	}
}
