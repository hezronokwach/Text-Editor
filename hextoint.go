package reloaded

import (
	"fmt"
	"strconv"
)

// HexToInt function converts a hexadecimal string to its decimal representation.
// It takes a hexadecimal string as input and returns its decimal representation as a string.
// If an error occurs during conversion, it prints an error message.
func HexToInt(hex string) string {
	number, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		fmt.Println("Error converting to hex")
	}
	return fmt.Sprint(number)
}
