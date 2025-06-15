package main

import (
	"fmt"
	"math/cmplx"
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

	// 2. Floating point (decimals)
	var i float32 = 3.14              // Single precision
	var j float64 = 3.141592653589793 // Double precision (more accurate)

	// 3. Complex numbers  (for math/science applications)
	var k complex64 = complex(1, 2) // 1+2i
	var l complex128 = cmplx.Sqrt(-5 + 12i)

	// 4. BYTE and RUNE
	var m byte = 65  // byte is an alias for uint8 (ASCII character)
	var n rune = 'あ' // rune is an alias for int32 (Unicode character)

	// 5. Boolean
	var o bool = true
	var p bool = false

	// 6. String
	var q string = "Hello, Go!"

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

	fmt.Println("\nFLOATING POINT:")
	fmt.Println("float32:", i)
	fmt.Println("float64:", j)

	fmt.Println("\nCOMPLEX NUMBERS:")
	fmt.Println("complex64:", k)
	fmt.Println("complex128:", l)

	fmt.Println("\nBYTE and RUNE:")
	fmt.Println("byte:", m, "as character:", string(m))
	fmt.Println("rune:", n, "as character:", string(n))

	fmt.Println("\nBOOLEAN:")
	fmt.Println("bool (true):", o)
	fmt.Println("bool (false):", p)

	fmt.Println("\nSTRING:")
	fmt.Println("string:", q)
}
