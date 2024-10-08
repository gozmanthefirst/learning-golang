package main

import "fmt"

// Variables
func variables() {
	var num1 = 4 // type inferred

	// Explicitly typed varibales
	var num2 int      // `int` variables have a default value of 0
	var num3 float64  // `float64` variables have a default value of 0
	var name string   // `string` varibales have a default values of ""
	var isActive bool // `bool` varibales have a default values of false

	// It is also possible to explicitly declare a variable and initialize a value at the same time
	var dollarToNaira float64 = 1667.23

	// Short variable declaration operator
	// Can't be used outside a function body
	surname := "Sunday"

	// Declaring and initializing multiple variables at once
	firstName, lastName, age := "Chiagoziem", surname, 23

	// Also possible using the `var keyword
	var village, town = "Ururo", "Umunze"

	// Alternate way
	var (
		state   = "Anambra"
		country = "Nigeria"
	)

	// Decalring never-changing constants
	// A `const` variable can appear anywhere a `var` statement can.
	const planet = "Earth"

	// Unused variables
	notUsed := "This was not used"
	_ = notUsed // assinging the unused variable to a blank identifier fixes the `not used` error

	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(num3)
	fmt.Println(name)
	fmt.Println(isActive)
	fmt.Println(dollarToNaira)
	fmt.Println(firstName, lastName, age)
	fmt.Println(village, town)
	fmt.Println(state, country)
	fmt.Println(planet)
}

// Strings
func strings() {

}
