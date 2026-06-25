package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var ans int
	// Split the input by newlines
	inputList := strings.SplitSeq(input, "\n")
	for v := range inputList {
		// iterate over each row to calculate the digitalRoot
		n := digitalRoot(v)
		ans += n
	}

	return ans
}

func digitalRoot(input string) int {
	if len(input) <= 1 {
		n, _ := strconv.ParseInt(input, 10, 0)
		return int(n)
	}

	var ans int
	for _, v := range input {
		n, _ := strconv.ParseInt(string(v), 10, 0)
		ans += int(n)
	}

	return digitalRoot(strconv.Itoa(ans))
}

// solver selects which implementation runs. -solver=ref uses the A-grade
// reference (reference.go); the default uses my solution above. The test suite
// reads the same flag, so `go test -solver=ref` runs every unit test through
// the reference instead.
var solver = flag.String("solver", "mine", `which solution to run: "mine" or "ref"`)

func digitalRootSumActive(input string) int {
	if *solver == "ref" {
		return digitalRootSumRef(input)
	}
	return digitalRootSum(input)
}

func main() {
	flag.Parse()
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error reading input.txt:", err)
		os.Exit(1)
	}
	fmt.Println(digitalRootSumActive(string(data)))
}
