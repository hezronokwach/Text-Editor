package main

import (
	"os"
	"testing"
)
func TestMainFunction(t *testing.T) {
	// Create a temporary input file
	inputFile := "input.txt"
	defer os.Remove(inputFile)

	// Write some sample input to the temporary input file
	inputText := "I am exactly how they (up, 2) describe me am me : ' awesome '"
	err := os.WriteFile(inputFile, []byte(inputText), 0644)
	if err != nil {
		t.Fatalf("Failed to create input file: %v", err)
	}

	// Define the expected output
	expectedOutput := "I am exactly How They describe me am me: 'awesome'"

	// Create a temporary output file
	outputFile := "output.txt"
	defer os.Remove(outputFile)

	// Call the main function with input and output files
	os.Args = []string{"cmd", inputFile, outputFile}
	main()

	// Read the content of the output file
	outputContent, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	// Check if the output matches the expected output
	if string(outputContent) != expectedOutput {
		t.Errorf("Output does not match expected output:\nExpected: %s\nActual: %s", expectedOutput, (outputContent))
	}
}
