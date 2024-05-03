package reloaded

import (
	"testing"
)
// Struct to define test cases for the Capitalize function
// arg1 means argument 1 and the expected stands for the 'result expected'
type capitalizeTest struct {
	arg1, expected string
}

// Slice of test cases for the Capitalize function
var capitalizeTests = []capitalizeTest{
	{"HEllo mr", "Hello Mr"},
	{"HEllo Mr", "Hello Mr"},
	{"hEllo", "Hello"},
}

// Test function to verify the correctness of the Capitalize function
func TestCapitalize(t *testing.T) {
	// Iterate over each test case
	for _, test := range capitalizeTests {
		// Call the Capitalize function with the test input
		output := Capitalize(test.arg1)
		// Check if the output matches the expected result
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}

// Benchmark function to measure the performance of the Capitalize function
func BenchmarkCapitalize(b *testing.B) {
	// Execute the Capitalize function repeatedly for benchmarking
	for i := 0; i < b.N; i++ {
		Capitalize("my NamE is Okwach")
	}
}
