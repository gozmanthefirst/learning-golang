package main

import (
	"fmt"
	"slices"
)

func main() {
	//* ARRAYS
	// For some reasons, arrays are rarely used directly in Go. All elements in an array must have a specified type.
	// Here are several ways to declare an array:

	// 1. Specify the size of the array, and the type of the elements in the array.

	var arr1 [3]int // This creates an array with 3 elements of type `int`

	// Since no value was specified, all 3 elements are given the zero value of `int`.
	// To give them initial values, we do something like this:

	var arr2 = [4]int{1, 2, 3, 4}

	// In a sparse array (most elements have a zero value), you can indicate the element that will have a non-zero value using their indices.
	var arr3 = [10]int{2, 6: 283, 7: 352, 99}

	// When creating an array, you can also replace the size of the array with `...`. Here, the size of the array will be gotten from the number of initial values passed to the array.
	var arr4 = [...]float64{12.34, 5, 67.89}

	// Arrays are compared using `==` and `!=`. two arrays are equal if they have the same size and values.

	// Multidimensional arrays (matrices), can be simulated using something like this:

	var arr5 [2][4]int // This depicts an array with two elements, each being an array with 4 elements.

	fmt.Println(arr1, arr2, arr3, arr4, arr5)

	// Arrays in Go are also written and read using the bracket notation.
	arr4[1] = 5.43
	fmt.Println(arr4, arr3[6])

	// To get the length of an array, we can use the built in len() function
	fmt.Println(len(arr4))

	// Arrays are difficult to work with in Go, and therefore are rarely used. The reasons for this include:
	// 1. The size of the array is considered to be a part of the array's type. Therefore, an array with type [3]int is a different type from an array with type [5]int.
	// 2. The types must also be resolved at compile time, therefore vartiables cannot be used to specify the type of the array.
	// 3. Array of different sizes cannot be converted into the same type using type conversion. Due to this, you can't write a function that can work with an array of any size, and you can't assign arrays of different sizes to the same variables.

	// The reason arrays exist in Go despite being a very limited feature is to provide the backign store for slices.

	//* SLICES
	// Slices are identical to arrays in many ways but there are some very important differences. They include:
	// 1. The size (either a number or `...`) is not specified when declaring a slice.

	var sli1 []int

	fmt.Println(sli1)

	// 2. The zero value for a slice is `nil`. A nil slice contains nothing.
	// 3. Two slices cannot be compared to each other using `==` and `!=`. They can only be compared to `nil`. Although there are functions from the `slices` pkg from the standard library that can be used to compare two slices.

	sli2 := []int{1, 2, 3, 4, 5}
	sli3 := []int{1, 2, 3, 4, 5}
	sli4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slices.Equal(sli2, sli3))
	fmt.Println(slices.Equal(sli2, sli4))

    //* `slices` FUNCTIONS
}
