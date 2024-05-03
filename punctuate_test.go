package reloaded

import (
	"testing"
)

// TestPunctuate function tests the Punctuate function with different input scenarios.
func TestPunctuate(t *testing.T) {
	// Define test cases with input and expected output
	tests := []struct {
		name     string     // Name of the test case
		input    []string   // Input slice of strings
		expected []string   // Expected output slice of strings
	}{		
		{name: "Single punctuation enclosing a word", input: []string{"This", "'", "text", "'", "is", "an", "example."}, expected: []string{"This", "'text'", "is", "an", "example."}},
		{name: "Punctuation mark at the end", input: []string{"This", "is", "an", "example", "Another", "sentence", "."}, expected: []string{"This", "is", "an", "example", "Another", "sentence."}},
		{name: "Apostrophe in the first position", input: []string{"'This", "an", "example.'"}, expected: []string{"'This", "an", "example.'"}},
		{name: "Apostrophe in a contraction", input: []string{"There's", "an", "example."}, expected: []string{"There's", "an", "example."}},
		{name: "Multiple punctuation marks", input: []string{"This", "is", "!!!", "an", "example."}, expected: []string{"This", "is!!!", "an", "example."}},
	}

	// Iterate over each test case
	for _, tc := range tests {
		// Run sub-tests for each test case
		t.Run(tc.name, func(t *testing.T) {
			// Call the Punctuate function with the test input
			got := Punctuate(tc.input)
			// Check if the output matches the expected result
			if !stringSliceEqual(got, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, got)
			}
		})
	}
}

// BenchmarkPunctuate function benchmarks the performance of the Punctuate function.
func BenchmarkPunctuate(b *testing.B) {
	// Execute the Punctuate function repeatedly for benchmarking
	for i := 0; i < b.N; i++ {
		slice := []string{"This", "'", "text", "'", "is", "an", "example."}
		Punctuate(slice)
	}
}

// stringSliceEqual function compares two slices of strings and returns true if they are equal.
func stringSliceEqual(a, b []string) bool {
	// Check if the length of the slices is equal
	if len(a) != len(b) {
		return false
	}
	// Iterate over each element in the slices
	for i, v := range a {
		// If any element differs, return false
		if v != b[i] {
			return false
		}
	}
	// If all elements are equal, return true
	return true
}
