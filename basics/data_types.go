package main

import (
	"fmt"
)

func main() {
	// Basic types in go

	// 1. Integers
	// Signed integers (can be positive or negative)
	var a int8 = 127                  // -128 to 127
	var b int16 = 32767               // -32768 to 32767
	var c int32 = 2147483647          // -2147483648 to 2147483647
	var d int64 = 9223372036854775807 // Very large range!
	var e int = 42                    // Size depends on system (32 or 64 bits)

	// Unsigned integers (only positive values)
	var f uint8 = 255    // 0 to 255
	var g uint16 = 65535 // 0 to 65535
	var h uint = 100     // Size depends on systen

	// Print everything to see the values
	fmt.Println("INTEGERS:")
	fmt.Println("int8:", a)
	fmt.Println("int16:", b)
	fmt.Println("int32:", c)
	fmt.Println("int64:", d)
	fmt.Println("int:", e)
	fmt.Println("uint8:", f)
	fmt.Println("uint16:", g)
	fmt.Println("uint:", h)
}
