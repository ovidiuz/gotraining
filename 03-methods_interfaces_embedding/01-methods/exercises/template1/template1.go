// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/Rj0QfwVPhX

// Declare a struct that represents a baseball player. Include name, atBats and hits.
// Declare a method that calculates a players batting average. The formula is Hits / AtBats.
// Declare a slice of this type and initialize the slice with several players. Iterate over
// the slice displaying the players name and batting average.
package main

import "fmt"

// Add imports.

// Declare a struct that represents a ball player.
// Include field called name, atBats and hits.

type ballPlayer struct {
	name string
	atBats int
	hits int
}

// Declare a method that calculates the batting average for a batter.
func (b *ballPlayer) average() float32 {
	return float32(b.atBats) / float32(b.hits)
}

// main is the entry point for the application.
func main() {
	// Create a slice of players and populate each player
	// with field values.
	var players []ballPlayer = []ballPlayer {
		ballPlayer{
			name:   "John",
			atBats: 10,
			hits:   3,
		},
		ballPlayer{
			name:   "Bill",
			atBats: 2,
			hits:   1,
		},
	}

	// Display the batting average for each player in the slice.
	for _, player := range players {
		average := player.average()
		fmt.Println(average)
	}
}
