package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	//* Literals
	// A Go literal can be an explicitly specified number, string or character (rune).
	// Literals are considered 'untyped'.

	//* Booleans
	// Only 'true' or 'false'. Zero value is 'false'.

	//* NUMERIC TYPES -> INTEGER TYPES
	// The types include, int8 to int64, and uint8 to uint64, in multiples of 2. unsigned contains only positive numbers.
	// Special integer types include,
	// - 'byte' -> which is basically the same thing as uint8.
	// - 'int' -> can be int32 or int64 depending on the CPU of the computer. This is the default type for integer literals.
	// - 'uint' -> can be uint32 or uint64 depending on the CPU of the computer, so basically the same thing as int, but the values are always 0 or positive.
	// Other special types include 'rune' and 'uintptr'. Will be learnt later.
	// Zero value for integer types is 0.
	// Integers can undergo the normal arithmetic operations. NOTE: The result of an integer division is always an integer.

	//* NUMERIC TYPES -> FLOATING POINT TYPES
	// There are only two — float32 and float64. The zero value is 0.
	// All arithmetic operators except for '%' can be used with floats.
	//! Do not use floating point types to represent money or any value that must have an exact decimal representation!
	//! Do not compare two floats with the usual comparison operators, if you can help it.
	// Go also has first-class support for complex numbers.

	// Complex numbers operation
	a := complex(2.5, 3.1)
	b := complex(10.2, 2)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(real(a))
	fmt.Println(imag(a))
	fmt.Println(cmplx.Abs(a))

	//* STRINGS AND RUNES
	// The zero value for a string is the empty string, `""`.
	// The rune type signifies a single code point, ie. a single character.

	//* EXPLICIT TYPE CONVERSION
	// Unlike most languages out there, Go doesn't support automatic type promotion/conversion, so as to avoid complications.
	// Even different-sized ints and floats must be converted to the same type to interact.

	// Type conversion between different types
	var c int = 10
	var d float64 = 30.2
	var sum1 float64 = float64(c) + d
	var sum2 int = c + int(d)
	fmt.Println(sum1, sum2)

	// Type conversion between different sizes of the same type
	var e int = 10
	var f byte = 100
	var sum3 int = e + int(f)
	var sum4 byte = byte(e) + f
	fmt.Println(sum3, sum4)

	//! One implication of the explicit type conversion is that, since there can't be implicit type conversion between types, there's no such thing as a 'truthy' or 'falsy' value.

	//* WAYS TO DECLARE VARIABLES
	// 1. Using the `var` keyword, name of variable, type of variable and the value.
	var g int = 34

	// 2. Using the `var` keyword, name of variable and the value. The type is inferred.
	var h = 100.12

	// 3. Using the `var` keyword, the name of the variable and the type. The variable is assigned the zero value of the type.
	var i string

	// 4. Declaring multiple variables of the same type at once with `var`, and assigning values to them.
	var j, k int = 20, 1212

	// 5. Declaring multiple variables of the same type, with them being given the zero value of the type.
	var l, m int

	// 6. Declaring multiple variables of different types, and assigning values to them.
	var n, o = 12.34, "wtf"

	// 7. Declaring multiple variables at once using a declaration list.
	var (
		p    int = 12
		q        = "wagwan!"
		r    bool
		s, t = 94, "hello"
	)

	// 8. Short declaration and assignment format.
	u, v := 74, "sigh"

	// A trick with using the `:=` is that as long as you're declaring a new variable, you can reassign an existing value
	v, w := "lfg!", 43.21

	fmt.Println(g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w)
}
