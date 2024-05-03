package reloaded

import (
	"fmt"
	"strconv"
)

// Binary function converts a binary string to its decimal representation.
// It takes a binary string as input and returns its decimal representation as a string.
// If an error occurs during conversion, it prints an error message.
func Binary(bin string) string {
	number, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		fmt.Println("Error converting to binary")
	}
	return fmt.Sprint(number)
}


