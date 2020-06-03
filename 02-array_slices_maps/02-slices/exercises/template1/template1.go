// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/mPKmyGNR2L

// Declare a nil slice of integers. Declare a nil slice of integers. Create a
// loop that appends 10 values to the slice. Iterate over the slice and display
// each value. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare a nil slice of integers.
	var arr []int

	// Appends numbers to the slice.
	for i := 0; i < 5; i++ {
		arr = append(arr, 5 - i)
	}

	// Display each value in the slice.
	for _, number := range arr {
		fmt.Println(number)
	}

	// Declare a slice of strings and populate the slice with names.
	var names []string = []string{"Joe", "Bill", "Todd", "Andrew", "Josh"}

	// Display each index position and slice value.
	for i, name := range names {
		fmt.Println(i, name)
	}

	// // Take a slice of index 1 and 2 of the slice of strings.
	var slice = names[1:3]

	// Display each index position and slice values for the new slice.
	for i, name := range slice {
		fmt.Println(i, name)
	}
}
