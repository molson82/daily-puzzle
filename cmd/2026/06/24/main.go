package main

import (
	"fmt"
	"os"
)

// Digital Root Sum
//
// Input: one non-negative integer per line (blank lines are ignored).
//
// The digital root of a number is found by repeatedly summing its digits
// until a single digit remains. For example:
//
//	942 -> 9+4+2 = 15 -> 1+5 = 6
//
// Compute the digital root of every number in the input and return the sum
// of all those digital roots.
//
// digitalRootSum returns the total described above.
func digitalRootSum(input string) int {
	// TODO: implement
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input.txt:", err)
		os.Exit(1)
	}
	fmt.Println(digitalRootSum(string(data)))
}
