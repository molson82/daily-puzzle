package main

import (
	"os"
	"testing"
)

func TestHeaviestShelf(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"single shelf", "1\n2\n3", 6},
		{"two shelves", "1\n2\n\n10", 10},
		{"sample", "10\n20\n\n5\n\n7\n8\n9", 30},
		{"leading/trailing blanks", "\n\n5\n\n\n6\n\n", 6},
		{"all zeros", "0\n0\n\n0", 0},
		{"first is heaviest", "100\n\n1\n2\n3", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heaviestShelf(tt.input); got != tt.want {
				t.Errorf("heaviestShelf(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestTopThreeShelves(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty", "", 0},
		{"single shelf", "1\n2\n3", 6},
		{"two shelves sums both", "1\n2\n\n10", 13},
		{"sample", "10\n20\n\n5\n\n7\n8\n9", 59},
		{"four shelves takes top three", "1\n\n2\n\n3\n\n4", 9},
		{"leading/trailing blanks", "\n\n5\n\n\n6\n\n", 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topThreeShelves(tt.input); got != tt.want {
				t.Errorf("topThreeShelves(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestHeaviestShelfSampleFile(t *testing.T) {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("reading sample.txt: %v", err)
	}
	if got, want := heaviestShelf(string(data)), 30; got != want {
		t.Errorf("heaviestShelf(sample.txt) = %d, want %d", got, want)
	}
	if got, want := topThreeShelves(string(data)), 59; got != want {
		t.Errorf("topThreeShelves(sample.txt) = %d, want %d", got, want)
	}
}
