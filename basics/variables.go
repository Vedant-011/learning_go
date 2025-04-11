package main

import "fmt"

func main() {
	// 1. Declaring variables with var
	var name string = "Vedant"
	var age int = 21
	var isGamer bool = true

	// 2. Type inference (Go figures out the type)
	var country = "India"
	var year = 2025

	// 3. Short variable declaration (only inside functions)
	hobby := "Breathing"
	score := 95.5
	kdRatio := 2.35 // float64 variable 

	// 4. Zero values (default values when you don't assign anything)
	var emptyInt int
	var emptyString string
	var emptyBool bool

	// Printing all values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Gamer:", isGamer)
	fmt.Println("Country:", country)
	fmt.Println("Year:", year)
	fmt.Println("Hobby:", hobby)
	fmt.Println("Score:", score)
	fmt.Println("KD Ratio:", kdRatio)

	fmt.Println("Empty Int:", emptyInt)
	fmt.Println("Empty String:", emptyString)
	fmt.Println("Empty Bool:", emptyBool)
}


/*
Name: Vedant
Age: 21
Is Gamer: true
Country: India
Year: 2025
Hobby: Breathing
Score: 95.5
KD Ratio: 2.35
Empty Int: 0
Empty String:
Empty Bool: false
*/