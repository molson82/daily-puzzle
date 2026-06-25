package main

import (
	"fmt"
	"os"
)

// Balanced Brackets
//
// Input: one expression per line. A line may mix the bracket characters
// ()[]{} with any other characters (letters, digits, spaces); non-bracket
// characters are ignored. Blank lines are skipped.
//
// A line is balanced if every opening bracket is closed by the matching
// type in the correct order: "([])" is balanced, "([)]" is not.
//
// countBalanced returns how many lines are balanced.
func countBalanced(input string) int {
	// TODO: implement
	return 0
}

// isBalanced reports whether s has properly matched, properly nested
// brackets, ignoring any non-bracket characters.
func isBalanced(s string) bool {
	// TODO: implement
	return false
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input.txt:", err)
		os.Exit(1)
	}
	fmt.Println(countBalanced(string(data)))
}
