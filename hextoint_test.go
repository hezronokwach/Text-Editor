package reloaded

import (
	"testing"
)

// Struct to define test cases for the HexToInt function
// arg1 means argument 1 and the expected stands for the 'result expected'
type hexTest struct {
	arg1, expected string
}

// Slice of test cases for the HexToInt function
var hexTests = []hexTest{
	{"1E", "30"},
	{"2B", "43"},
	{"5A", "90"},
}

// Test function to verify the correctness of the HexToInt function
func TestHex(t *testing.T) {
	// Iterate over each test case
	for _, test := range hexTests {
		// Call the HexToInt function with the test input
		output := HexToInt(test.arg1)
		// Check if the output matches the expected result
		if output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
// Benchmark function to measure the performance of the HexToInt function
func BenchmarkHex(b *testing.B) {
	// Execute the HexToInt function repeatedly for benchmarking
	for i := 0; i < b.N; i++ {
		HexToInt("10")
	}
}
