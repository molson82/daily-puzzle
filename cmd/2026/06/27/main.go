package main

import (
	"fmt"
	"os"
)

// Warehouse Shelves
//
// Input: groups of box weights, one non-negative integer per line. A blank
// line separates one shelf from the next. Leading/trailing and repeated blank
// lines should be tolerated.
//
// heaviestShelf returns the total weight of the heaviest shelf (the group of
// weights with the largest sum). Empty input -> 0.
//
// topThreeShelves returns the combined weight of the three heaviest shelves.
// With fewer than three shelves, sum all of them. Empty input -> 0.
func heaviestShelf(input string) int {
	// TODO: implement
	return 0
}

func topThreeShelves(input string) int {
	// TODO: implement
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input.txt:", err)
		os.Exit(1)
	}
	fmt.Println(heaviestShelf(string(data)))
	fmt.Println(topThreeShelves(string(data)))
}
