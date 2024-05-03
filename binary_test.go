package reloaded

import (
	"testing"
)

// Struct to define test cases for the Binary function
// arg1 means argument 1 and the expected stands for the 'result expected'
type binaryTest struct {
	arg1, expected string
}

// Slice of test cases for the Binary function
var binaryTests = []binaryTest{
	{"10", "2"},
	{"11", "3"},
	{"100", "4"},
}

// Test function to verify the correctness of the Binary function
func TestBinary(t *testing.T) {
	// Iterate over each test case
	for _, test := range binaryTests {
		// Call the Binary function with the test input
		output := Binary(test.arg1)
		// Check if the output matches the expected result
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
// Benchmark function to measure the performance of the Binary function
func BenchmarkBinary(b *testing.B) {
	// Execute the Binary function repeatedly for benchmarking
	for i := 0; i < b.N; i++ {
		Binary("10")
	}
}
