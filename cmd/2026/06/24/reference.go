package main

import "strings"

// digitalRootSumRef is a reference A-grade solution to Digital Root Sum.
//
// Improvements over the original solution (see the 2026-06-24 grade in
// PROGRESS.md):
//   - No strconv per digit: a digit rune's value is just r - '0' — no
//     allocation, no error to discard, and non-digits are skipped explicitly
//     instead of relying on ParseInt(" ") quietly returning 0.
//   - No string round-trip in the recursion: the digital root comes from a
//     closed form, so we never rebuild a string to recurse.
//   - Overflow-proof for arbitrarily large numbers: a number is congruent to
//     its digit sum mod 9, so we reduce by summing digit characters and never
//     parse the full (possibly huge) number into an int.
func digitalRootSumRef(input string) int {
	total := 0
	for line := range strings.SplitSeq(input, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		// Sum the digit characters. digitSum(n) ≡ n (mod 9), so the digital
		// root of this (small) sum equals the digital root of n itself.
		sum := 0
		for _, r := range line {
			if r >= '0' && r <= '9' {
				sum += int(r - '0')
			}
		}
		total += digitalRootRef(sum)
	}
	return total
}

// digitalRootRef returns the digital root of a non-negative integer in O(1).
// For positive n the digital root cycles 1..9, and it is 0 only for 0:
//
//	dr(n) = 0           if n == 0
//	        1+(n-1)%9   otherwise
func digitalRootRef(n int) int {
	if n == 0 {
		return 0
	}
	return 1 + (n-1)%9
}
