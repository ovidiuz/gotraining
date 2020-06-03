// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/-JBSUoux-v

// Declare and make a map of integer values with a string as the key. Populate the
// map with five values and iterate over the map to display the key/value pairs.
package main

import "fmt"

// Add imports.

// main is the entry point for the application.
func main() {
	// Declare and make a map of integer type values.
	var salary map[string]int = make(map[string]int)

	// Initialize some data into the map.
	salary["John"] = 2000
	salary["Andrew"] = 2500
	salary["Kyle"] = 1800


	// Display each key/value pair.
	for key, value := range salary {
		fmt.Println(key, value)
	}
}
