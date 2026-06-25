package main

import (
	"os"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{"empty", "", true},
		{"simple pair", "()", true},
		{"nested mixed", "([])", true},
		{"sequence", "()[]{}", true},
		{"interleaved", "([)]", false},
		{"unclosed", "(((", false},
		{"unopened", ")(", false},
		{"only closer", "]", false},
		{"ignore letters", "a(b)c[d]{e}", true},
		{"no brackets", "hello world", true},
		{"deep nest", "{[()()]}", true},
		{"wrong close", "(]", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.in); got != tt.want {
				t.Errorf("isBalanced(%q) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCountBalanced(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"empty input", "", 0},
		{"sample", "()\n([)]\n(a[b]{c})", 2},
		{"blank lines skipped", "()\n\n[]\n", 2},
		{"all bad", "(\n)\n][", 0},
		{"plain text counts", "no brackets here\n(ok)", 2},
		{"trailing newline", "{}\n", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBalanced(tt.input); got != tt.want {
				t.Errorf("countBalanced(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestCountBalancedSampleFile(t *testing.T) {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("reading sample.txt: %v", err)
	}
	if got, want := countBalanced(string(data)), 2; got != want {
		t.Errorf("countBalanced(sample.txt) = %d, want %d", got, want)
	}
}
