// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/c9Qrsq8QFe

// Copy the code from the template. Declare a new type called hockey
// which embeds the sports type. Implement the matcher interface for hockey.
// When implementing the Search method for hockey, call into the Search method
// for the embedded sport type to check the embedded fields first. Then create
// two hockey values inside the slice of matchers and perform the search.
package main

import (
	"fmt"
	"strings"
)

// matcher defines the behavior required for performing searches.
type matcher interface {
	Search(searchTerm string) bool
}

// sport represents a sports team.
type sport struct {
	team string
	city string
}

// Search checks the value for the specified term.
func (s sport) Search(searchTerm string) bool {
	if strings.Contains(s.team, searchTerm) ||
		strings.Contains(s.city, searchTerm) {
		return true
	}

	return false
}

// Declare a struct type named hockey that represents specific
// hockey information. Have it embed the sport type first.
type hockey struct {
	sport
	iceRink string
}

// Implement the matcher interface for hockey.
func (h hockey) Search(searchTerm string) bool {
	// Make sure you call into Search method for the embedded
	// sport type.
	return h.sport.Search(searchTerm) || strings.Contains(h.iceRink, searchTerm)
}

// main is the entry point for the application.
func main() {
	// Define the term to search.
	searchTerm := "Moscow"

	// Create a slice of matcher values to search.
	matchers := []matcher{
		hockey{sport{"Panthers", "Moscow"}, "USA"},
		hockey{sport{"Canadians", "Montreal"}, "Canada"},
		hockey{sport{"Canadians", "Montreal"}, "Moscow Russia"},
	}

	// Display what we are searching for.
	fmt.Println("Searching For:", searchTerm)

	// Range of each matcher value and check the search term.
	for _, m := range matchers {
		if m.Search(searchTerm) {
			fmt.Printf("FOUND: %+v", m)
		}
	}
}
