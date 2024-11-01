package main

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
	"unicode/utf8"
)

// * Variables
func variables() {
	var num1 = 4 // type inferred

	// Explicitly typed variables.

	var num2 int      // `int` variables have a default value of 0.
	var num3 float64  // `float64` variables have a default value of 0
	var name string   // `string` variables have a default values of ""
	var isActive bool // `bool` variables have a default values of false

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

	// Declaring never-changing constants.
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

	fmt.Println("")
}

// * Strings
func strings() {
	var firstName string = "Erikon"
	fmt.Println(firstName)

	// Special Characters
	address := "02, Ehosi Square,\nUmuarori, Imeohia Kindred,\nUruro, Umunze,\nAnambra state, Nigeria."
	fmt.Println(address)

	// Raw Strings
	quotation := `"Some shit no one 
    ever said."
    -- No One`
	fmt.Println(quotation)

	// Unicode Characters

	str1 := "你好,世界"   // Chinese
	str2 := "こんにちは世界" // Japanese

	fmt.Println(str1)
	fmt.Println(str2)

	// Each chinese or japanese character takes up 3 bytes so uisng `len()` will give us strange results
	fmt.Println(len(str2)) // 12

	// To get the accurate number of characters, use utf8.RuneCountInString()
	fmt.Println(utf8.RuneCountInString(str2))

	fmt.Println("")
}

// * Performing Type Conversions
func typeConverions() {
	// The type of a variable can be gotten using two methods
	start := time.Now()

	// Method 1
	// Using `%T` printing verb in the Printf() function
	fmt.Printf("%T\n", start)

	// Method 2
	// Using the reflect pkg
	fmt.Println(reflect.TypeOf(start))
	fmt.Println(reflect.ValueOf(start).Kind())

	// var age int
	// fmt.Print("Please enter your age: ")
	// fmt.Scanf("%d", &age)
	// fmt.Println("You entered:", age)

	// To make the program more robust, you might want to change the input variable to a string before assigning the input value.

	var input string
	fmt.Print("Please enter your age: ")
	fmt.Scanf("%s", &input)
	fmt.Println("You entered:", input)

	// After the input is read, we can then use the strconv pkg's Atoi() to try to convert the string to an integer value.
	//? Atoi means ASCII to integer
	age, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Your age is:", age)
	}

	// Converting strings to other types
	// To booleans
	b, err := strconv.ParseBool("t") // t for true, f for false
	fmt.Println(err)
	fmt.Println(b)

	// To float
	f, err := strconv.ParseFloat("3.142", 64) // 64-bit float
	fmt.Println(err)
	fmt.Println(f)

	// To integer
	i, err := strconv.ParseInt("-10.23", 10, 64) // base 10, 64-bit integer.
	// An error will be returned since -10.23 is not an integer
	fmt.Println(err)
	fmt.Println(i)

	// To convert between the various numeric data types, we can use int(), float64(), float32()
	int1 := 12
	float1 := float64(int1)
	float2 := float32(float1)
	int2 := int(float2)

	fmt.Println(int2)

	// Interpolating strings
	queue := 32
	name := "Gozman"

	// To interfpolate, we can use fmt.Sprintf()
	queueString := fmt.Sprintf("Hello %s, your queue number is %d.", name, queue)
	fmt.Println(queueString)
}
