package main

import (
	"fmt"
	"reflect"
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

	// It is impossible to read or write past the end of an array, or to use negative index.

	// To get the length of an array, we can use the built in len() function
	fmt.Println(len(arr4))

	// Arrays are difficult to work with in Go, and therefore are rarely used. The reasons for this include:
	// 1. The size of the array is considered to be a part of the array's type. Therefore, an array with type [3]int is a different type from an array with type [5]int.
	// 2. The types must also be resolved at compile time, therefore vartiables cannot be used to specify the type of the array.
	// 3. Array of different sizes cannot be converted into the same type using type conversion. Due to this, you can't write a function that can work with an array of any size, and you can't assign arrays of different sizes to the same variables.

	// The reason arrays exist in Go despite being a very limited feature is to provide the backing store for slices.

	//* SLICES
	// Slices are identical to arrays in many ways but there are some very important differences. They include:
	// 1. The size (either a number or `...`) is not specified when declaring a slice.

	var sl1 []int

	fmt.Println(sl1)

	// 2. The zero value for a slice is `nil`. A nil slice contains nothing.
	// 3. Two slices cannot be compared to each other using `==` and `!=`. Slices can only use `==` and `!=` to be compared to `nil`. However, there are functions from the `slices` package from the standard library that can be used to compare two slices.

	sl2 := []int{1, 2, 3, 4, 5}
	sl3 := []int{1, 2, 3, 4, 5}
	sl4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sl1 == nil)
	fmt.Println(slices.Equal(sl2, sl3))
	fmt.Println(slices.Equal(sl2, sl4))

	//* `slices` FUNCTIONS
	//* `len`
	// This returns the length of a slice, just as it did with arrays. Passing a `nil` slice to `len` will return 0.

	fmt.Println(len(sl1))
	fmt.Println(len(sl2))

	//* `append`
	// This is used to pass new elements to a slice. It takes at least 2 parameters, the slice being appended to, and the value that is being appended. The type of the elements of the slice must be the same with the type of the variable being appended to the slice. When using the `append` function, it is recommended that the slice that is returned should be assigned to the variable (slice) that was passed in to the `append` function (although, it can be assigned to any slice).

	sl1 = append(sl4, 12)
	fmt.Println(sl1)

	// More than one values can also be passed to be appended to a slice.
	sl2 = append(sl2, 132, 101)
	fmt.Println(sl2)

	//* Capacity
	// The capacity of a slice is the number of consecutive memory locations that has been reserved for the slice.

	fmt.Println(cap(sl1))
	fmt.Println(cap(sl2))

	//* `make`
	// Slices being able to grow automatically is a nice feature to have, but if the required number of values to be put in a slice is know before hand, it is a good practice to create it with the correct initial capacity. We can do that using `make`.
	// `make` allows for the creation of empty slices that have a specified length or capacity. It allows you to specify the type, length and optionally, the capacity of the slice.

	// This will create a slice containing 5 values with each one being equal to zero.
	sl5 := make([]int, 5)

	// Trying to append to `sl5` will add the new value to the end of the 5th value instead of making it the first value, since `append` always increases the length of a slice.
	sl5 = append(sl5, 1, 2) // This will result in [0, 0, 0, 0, 0, 1, 2]

	// We can also specify the capacity of the slice.
	sl6 := make([]float64, 10, 20) // here, 10 is the length and 20 is the capacity

	// We can assign a length of 0 to a slice to allow the elements of the slice to be added using `appended`.
	sl7 := make([]string, 0, 20)
	sl7 = append(sl7, "wagwan", "wtf!")

	fmt.Println(sl5)
	fmt.Println(sl6)
	fmt.Println(sl7)

	//* `clear`
	// Slices can be emptied using `clear`. This turns the values of the slice to their zero value. The length of the slice remains unchanged.

	sl8 := []int{43, 987, 478}
	fmt.Println(sl8, len(sl8))
	clear(sl8)
	fmt.Println(sl8, len(sl8))

	//* Declaring Slices
	// When creating a slice that will likely not grow at all, always favour nil slices to zero-length slices.

	// nil slice
	var sl9 []string

	// zero-length slice
	sl10 := []int{}

	fmt.Println(sl9, sl10)

	// If you have a good idea on how large the slice needs to be but do not know what those values will be, then use make.
	// The best way is to use it with a zero length and a specified capacity.

	sl11 := make([]string, 0, 100)
	fmt.Println(sl11)

	sl11 = append(sl11, "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z")

	//* Slicing Slices
	// A slice can be sliced (pun intended), by using the following notation:

	// s2 := s1[x:y]
	// where s1 is the slice being sliced, s2 is the subslice being created, x is the index of the element that the subslice will start from, and y is the index of the element that is one past the element that will be the last element in the subslice.

	sl12 := sl11[0:5]
	sl13 := sl11[3:]  // This will create a subslice from the element at index 3 of the old slice to the end of the old slice.
	sl14 := sl11[:11] // This will create a subslice from the beginning of the old slice to the element before the element at index 11 of the old slice.
	sl15 := sl11[:]   // This will create a subslice from the beginning of the old slice to the end of the old slice, which means it will create a subslice identical to the old slice.

	fmt.Println(sl12)
	fmt.Println(sl13)
	fmt.Println(sl14)
	fmt.Println(sl15)

	// Using append with subslices can be very tricky.
	slX := make([]string, 0, 5)
	slX = append(slX, "a", "b", "c", "d")
	slY := slX[:2]
	slZ := slX[2:]
	fmt.Println(cap(slX), cap(slY), cap(slZ))
	slY = append(slY, "i", "j", "k")
	slX = append(slX, "x")
	slZ = append(slZ, "y")
	fmt.Println("x:", slX)
	fmt.Println("y:", slY)
	fmt.Println("z:", slZ)

	// Therefore, it is adviced to either avoid using append with slices, or use a 'full slice expression'.
	// The 'full slice expression' includes a third part, which indicates the last position in the parent slice’s capacity that’s available for the subslice.

	slXX := make([]string, 0, 5)
	slXX = append(slXX, "a", "b", "c", "d")
	slYY := slXX[:2:2]
	slZZ := slXX[2:4:4]

	fmt.Println(slYY)
	fmt.Println(slZZ)

	// The third part indicates the last position of the original slice that is available to the subslice. The capacity of the subslice can then be calculated by subtracting the starting offset from the third part of the 'full slice expression'.

	//* `copy`
	// To create a slice from another slice but make it independent of the original slice, you can use the `copy` function. The copy function is used like this:

	// z := copy(x, y)
	// where x is the destination slice, y is the source slice, and z is the number of elements copied.
	// A subslice can also be the source slice.

	//* Converting an array to a slice
	// A slice can be created from an array or subset of an array by creating a subslice from the array.

	xArray := [5]int{1, 2, 3, 4, 5}
	xSlice := xArray[:]

	fmt.Println(reflect.TypeOf(xArray), reflect.TypeOf(xSlice)) // will return [5]int []int

	//* Converting a slice to an array
	// An array can be created from a slice by performing an explicit type conversion.

	ySlice := []int{1, 2, 3, 4, 5}
	yArray := [3]int(ySlice)

	fmt.Println(reflect.TypeOf(ySlice), reflect.TypeOf(yArray))

    //* Maps
}
