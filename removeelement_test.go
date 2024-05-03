package reloaded

import "testing"

// Struct to define test cases for the RemoveElement function
type removeElementTest struct {
	argSlice         []string
	argIndexToRemove int
	expectedSlice    []string
}

// Slice of test cases for the RemoveElement function
var removeElementTests = []removeElementTest{
	{
		argSlice:         []string{"a", "b", "c", "d", "e"},
		argIndexToRemove: 2,
		expectedSlice:    []string{"a", "b", "d", "e"},
	},
	{
		argSlice:         []string{"a", "b", "c", "d", "e"},
		argIndexToRemove: 0,
		expectedSlice:    []string{"b", "c", "d", "e"},
	},
	{
		argSlice:         []string{"a", "b", "c", "d", "e"},
		argIndexToRemove: 4,
		expectedSlice:    []string{"a", "b", "c", "d"},
	},
}

// Test function to verify the correctness of the RemoveElement function
func TestRemoveElement(t *testing.T) {
	// Iterate over each test case
	for _, test := range removeElementTests {
		// Make a copy of the input slice
		slice := make([]string, len(test.argSlice))
		copy(slice, test.argSlice)
		// Call the RemoveElement function with the test input
		RemoveElement(&slice, test.argIndexToRemove)
		// Check if the output matches the expected result
		if !slicesEqual(slice, test.expectedSlice) {
			t.Errorf("Output %v not equal to expected %v", slice, test.expectedSlice)
		}
	}
}

// Benchmark function to measure the performance of the RemoveElement function
func BenchmarkRemoveElement(b *testing.B) {
	// Execute the RemoveElement function repeatedly for benchmarking
	for i := 0; i < b.N; i++ {
		slice := []string{"a", "b", "c", "d", "e"}
		RemoveElement(&slice, 2)
	}
}

// slicesEqual function compares two slices of strings and returns true if they are equal.
func slicesEqual(a, b []string) bool {
	// Check if the length of the slices is equal
	if len(a) != len(b) {
		return false
	}
	// Iterate over each element in the slices
	for i := range a {
		// If any element differs, return false
		if a[i] != b[i] {
			return false
		}
	}
	// If all elements are equal, return true
	return true
}
