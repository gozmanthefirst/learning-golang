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

	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

    //* STRINGS AND RUNES
    // The zero value for a string is the empty string, '""'.
    // The rune type signifies a single code point, ie. a single character.
}
