package main

import "fmt"

// Exercises
func questionNum01() {
	var i = 20
	var f float64 = float64(i)

	fmt.Println(i, f)
}

func questionNum02() {
	const value = 8976
	var i int = value
	var f float64 = value

	fmt.Println(i, f)
}

func questionNum03() {
	var b byte = 255
	var smallI int32 = 2147483647
	var bigI uint64 = 18446744073709551615

	fmt.Println(b, smallI, bigI)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}
