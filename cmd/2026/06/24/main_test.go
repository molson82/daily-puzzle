package main

import (
	"os"
	"testing"
)

func TestDigitalRootSum(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"single digit", "5", 5},
		{"zero", "0", 0},
		{"single multi-digit", "942", 6},
		{"sample", "5\n10\n99", 15},
		{"with blank lines", "10\n\n11\n\n12\n", 6},
		{"nines and powers", "99\n100", 10},
		{"big number", "9999999999", 9},
		{"trailing whitespace", "  12  \n 345 ", 6},
		{"empty input", "", 0},
		{"multiples of nine", "9\n18\n27\n36", 36},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digitalRootSum(tt.input); got != tt.want {
				t.Errorf("digitalRootSum(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestDigitalRootSumSampleFile(t *testing.T) {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("reading sample.txt: %v", err)
	}
	if got, want := digitalRootSum(string(data)), 15; got != want {
		t.Errorf("digitalRootSum(sample.txt) = %d, want %d", got, want)
	}
}
