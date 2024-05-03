package reloaded

import (
	"testing"
)
// Struct to define test cases for the AddVowel function
type addVowelTest struct {
	name          string
	argSlice      []string
	expectedSlice []string
}

// Slice of test cases for the AddVowel function
var addVowelTests = []addVowelTest{
	{
		name:          "lowercase starting with a vowel",
		argSlice:      []string{"a", "umbrella"},
		expectedSlice: []string{"an", "umbrella"},
	},
	{
		name:          "uppercase starting with a vowel",
		argSlice:      []string{"A", "umbrella"},
		expectedSlice: []string{"An", "umbrella"},
	},
}

// Test function to verify the correctness of the AddVowel function
func TestAddVowel(t *testing.T) {
	// Iterate over each test case
	for _, test := range addVowelTests {		
		// Call the AddVowel function with the test input
		slice := AddVowel(test.argSlice)		
		// Check if the output matches the expected result
		if !slicesEqual(slice, test.expectedSlice) {
			t.Errorf("Output %v not equal to expected %v", slice, test.expectedSlice)
		}
	}
}

// Benchmark function to measure the performance of the AddVowel function
func BenchmarkAddVowel(b *testing.B) {
	// Execute the AddVowel function repeatedly for benchmarking
	for i := 0; i < b.N; i++ {
		AddVowel([]string{"a", "umbrella"})
	}
}
