package main

import "fmt"

// Define a struct outside the main function
// A struct is like a blueprint for a custom data type
type Person struct {
	FirstName string // Field names are capitalized by convention
	LastName  string
	Age       int
	Email     string
	Address   Address // Nested struct
}

// Another struct that will be nested
type Address struct {
	Street  string
	City    string
	Country string
	ZipCode string
}

func main() {
	// MAPS
	fmt.Println("=== MAPS ===")
	// Maps are Gos built-in hash table or dictionary type
	// They store key-value pairs

	// 1. DECLARING A MAP
	// Format: map[keyType]valueType
	var scores map[string]int // A nil map (not initialized yet)
	// You can't add to a nil map!

	// 2. INITIALIZING A MAP
	// Using make function
	scores = make(map[string]int)

	// 3. ADDING KEY-VALUE PAIRS
	scores["Saber"] = 91
	scores["Rem"] = 87
	scores["Holo"] = 95

	// 4. MAP LITERALS (declaring and initializing at once)
	ages := map[string]int{
		"Saber": 24,
		"Rem":   18,
		"Holo":  20,
	}

	// 5. ACCESSING VALUES
	fmt.Println("Rem's score:", scores["Rem"])
	fmt.Println("Saber's age:", ages["Saber"])

	// 6. CHECKING IF A KEY EXISTS
	holoAge, exists := ages["Holo"]
	if exists {
		fmt.Println("Holo's age is", holoAge)
	} else {
		fmt.Println("Holo's age is not known")
	}

	// So What about a non-existent key?
	vedantAge, exists := ages["Vedant"]
	fmt.Println("Vedant exists:", exists)         // false
	fmt.Println("Vedant's age value:", vedantAge) // 0 (zero value is returned when key doesn't exist)

	// 7. DELETING A KEY-VALUE PAIR
	delete(scores, "Rem")
	fmt.Println("After deleting Rem:", scores)

	// 8. MAP LENGTH
	fmt.Println("Number of people with scores:", len(scores))

	// 9. ITERATING OVER MAPS
	fmt.Println("\nAll ages:")
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}

	fmt.Println("\n=== STRUCTS ===")
	// Structs are collections of fields
	// They let you group related data together

	// 1. CREATING A STRUCT INSTANCE
	var saber Person
	// Fields have zero values by default
	fmt.Println("Empty person:", saber)

	// 2. ASSIGNING VALUES TO FIELDS
	saber.FirstName = "Saber"
	saber.LastName = "Pendragon"
	saber.Age = 24
	saber.Email = "hello@example.com"
	saber.Address.Street = "12 Marine Drive"
	saber.Address.City = "Mumbai"
	saber.Address.Country = "India"
	saber.Address.ZipCode = "400001"

	// 3. STRUCT LITERAL (declaring and initializing at once)
	rem := Person{
		FirstName: "Rem",
		LastName:  "Ayanami",
		Age:       18,
		Email:     "hello@example.com",
		Address: Address{
			Street:  "88 Carter Road",
			City:    "Mumbai",
			Country: "India",
			ZipCode: "400050",
		},
	}

	// 4. ABBREVIATED STRUCT LITERAL (not recommended, order matters)
	// This is less readable and prone to errors if struct definition changes
	holo := Person{"Holo", "wolf", 20, "hello@example.com", Address{"101 Lokhandwala Complex", "Mumbai", "India", "400053"}}

	// 5. ACCESSING STRUCT FIELDS
	fmt.Println("Saber's full name:", saber.FirstName, saber.LastName)
	fmt.Println("Rem's city:", rem.Address.City)

	// 6. COMPARING STRUCTS
	// Structs can be compared with == if all fields are comparable
	person1 := Person{FirstName: "Saber", LastName: "Pendragon"}
	person2 := Person{FirstName: "Saber", LastName: "Pendragon"}
	fmt.Println("Are person1 and person2 equal?", person1 == person2) // true

	// Print all people
	fmt.Println("\nPeople information:")
	fmt.Printf("Saber: %+v\n", saber) // %+v shows field names
	fmt.Printf("Rem: %+v\n", rem)
	fmt.Printf("Holo: %+v\n", holo)
}
