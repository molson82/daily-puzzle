package main

import (
	"fmt"
	"os"
)

// Run-Length Encoding
//
// Input: one string per line (ASCII). Blank lines are skipped.
//
// runLengthEncode compresses a string by replacing each run of the same
// character with its run length followed by the character:
//
//	"aaabbc"     -> "3a2b1c"
//	"wwwwwwwwww"  -> "10w"   (counts can be multi-digit)
//	""           -> ""
//
// A line is "compressible" when its encoding is strictly shorter than the
// original line. countCompressible returns how many non-blank lines are
// compressible.
func runLengthEncode(s string) string {
	// TODO: implement
	return ""
}

func countCompressible(input string) int {
	// TODO: implement
	return 0
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input.txt:", err)
		os.Exit(1)
	}
	fmt.Println(countCompressible(string(data)))
}
