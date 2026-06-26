package main

import (
	"os"
	"testing"
)

func TestRunLengthEncode(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"single", "a", "1a"},
		{"classic", "aaabbc", "3a2b1c"},
		{"no runs", "abc", "1a1b1c"},
		{"multi digit", "wwwwwwwwww", "10w"},
		{"mixed case distinct", "aAaA", "1a1A1a1A"},
		{"spaces count", "a   a", "1a3 1a"},
		{"long then short", "aaaaab", "5a1b"},
		{"trailing run", "abbbbb", "1a5b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runLengthEncode(tt.in); got != tt.want {
				t.Errorf("runLengthEncode(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestCountCompressible(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty input", "", 0},
		// "aaabbc"->"3a2b1c" (6==6, not shorter); "xy"->"1x1y" (longer);
		// "zzzz"->"4z" (shorter, compressible)
		{"sample", "aaabbc\nxy\nzzzz", 1},
		{"blank lines skipped", "aaaa\n\nbb\n", 1},
		{"multi digit run", "wwwwwwwwww", 1},
		{"none compress", "abc\nde\nf", 0},
		{"all compress", "aaa\nbbbb\nccccc", 3},
		{"trailing newline", "aaaa\n", 1},
		{"break even not counted", "aabb", 0}, // "aabb"->"2a2b" (4==4)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countCompressible(tt.input); got != tt.want {
				t.Errorf("countCompressible(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountCompressibleSampleFile(t *testing.T) {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("reading sample.txt: %v", err)
	}
	if got, want := countCompressible(string(data)), 2; got != want {
		t.Errorf("countCompressible(sample.txt) = %d, want %d", got, want)
	}
}
