package reloaded

// Capitalize function capitalizes the first letter of each word in a string.
// Non-alphabetic characters and digits do not affect capitalization.
func Capitalize(s string) string {
	var capitalString string
	capitalizerFlag := true // Flag to track whether the next character should be capitalized
	for _, value := range s {
		// Check if the character is alphabetic or a digit
		if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') || (value >= '0' && value <= '9') {
			// Capitalize the first character of a word
			if capitalizerFlag {
				if value >= 'a' && value <= 'z' {
					capitalString += string(value - 32) // Convert lowercase to uppercase
				} else {
					capitalString += string(value)
				}
				capitalizerFlag = false
			} else { // Keep the rest of the word in lowercase
				if value >= 'A' && value <= 'Z' {
					capitalString += string(value + 32) // Convert uppercase to lowercase
				} else {
					capitalString += string(value)
				}
			}
		} else { // Non-alphabetic characters or digits
			capitalString += string(value)
			capitalizerFlag = true // Reset the capitalizer flag for the next word
		}
	}
	return capitalString
}
