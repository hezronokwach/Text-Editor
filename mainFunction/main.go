package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"reloaded"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <sample.txt> <result.txt>")
		return
	}
	// Retrieve command-line arguments
	args := os.Args[1:]

	// Read input text from the file specified in the first argument
	inputText, error := os.ReadFile(args[0])
	if error != nil {
		fmt.Println("Error reading file")
	}
	if len(inputText) == 0 {
		fmt.Println("Empty file")
		return
	}
	// Split the input text into words based on spaces
	words := strings.Split(string(inputText), " ")

	for i, word := range words {
		// Switch statement to handle different transformation patterns
		switch word {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			// Remove the transformation pattern and its associated word from the slice
			reloaded.RemoveElement(&words, i)
		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			reloaded.RemoveElement(&words, i)
		case "(cap)":
			words[i-1] = reloaded.Capitalize(words[i-1])
			reloaded.RemoveElement(&words, i)
		case "(hex)":
			words[i-1] = reloaded.HexToInt(words[i-1])
			reloaded.RemoveElement(&words, i)
		case "(bin)":
			words[i-1] = reloaded.Binary(words[i-1])
			reloaded.RemoveElement(&words, i)
		case "(up,", "(low,", "(cap,":
			// Handle transformation patterns with a following number
			transformWithNumber(&words, i)
		}
	}
	// Punctuate the words
	final := reloaded.Punctuate(words)

	// Add vowels to the words
	reloaded.AddVowel(final)

	// Join the final words into a string separated by spaces
	finalString := strings.Join(final, " ") + "\n"

	// Write the final string to the file specified in the second argument
	err := os.WriteFile(args[1], []byte(finalString), 0o666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}

// Function to handle transformations with a number
func transformWithNumber(words *[]string, i int) {
	// Get the transformation type and number from the slice
	transformType := (*words)[i]
	b := strings.Trim((*words)[i+1], (*words)[i+1][len((*words)[i+1])-1:])
	// Convert the number to an integer
	number, err := strconv.Atoi(string(b))
	if number > len((*words)[:i]) {
		fmt.Println("Input Number equal to the length of the previous strings")
		return
	}
	if err != nil {
		fmt.Printf("Error converting number: %s\n", err)
		return
	}
	// Iterate over the specified number of words before the transformation pattern
	for j := 1; j <= number; j++ {
		// Apply the transformation to each word
		switch transformType {
		case "(up,":
			(*words)[i-j] = strings.ToUpper((*words)[i-j])
		case "(low,":
			(*words)[i-j] = strings.ToLower((*words)[i-j])
		case "(cap,":
			(*words)[i-j] = reloaded.Capitalize((*words)[i-j])
		}
	}
	// Remove the transformation pattern and its associated number from the slice
	reloaded.RemoveElement(words, i)
	reloaded.RemoveElement(words, i)
}
