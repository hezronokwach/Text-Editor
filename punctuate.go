package reloaded

func Punctuate(s []string) []string {
	punctuationMarks := []string{",", ".", "!", "?", ":", ";"}
	
	for i, word := range s {
		for _, mark := range punctuationMarks {
			// Check if the first and last characters of the word are the same punctuation mark
			// and the word is not the last word in the slice
			if string(word[0]) == mark && string(word[len(word)-1]) == mark && s[i] != s[len(s)-1] {
				// Concatenate the current word with the previous word and remove the current word from the slice
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	
	for i, word := range s {
		for _, mark := range punctuationMarks {
			// Check if the first character of the word is a punctuation mark
			// and the last character is not the same punctuation mark
			if string(word[0]) == mark && string(word[len(word)-1]) != mark {
				// Add the punctuation mark to the previous word and remove it from the current word
				s[i-1] += mark
				s[i] = word[1:]
			}
		}
	}
	
	for i, word := range s {
		for _, mark := range punctuationMarks {
			// Check if the first character of the word is a punctuation mark
			// and the word is the last word in the slice
			length := len(s)
			if (string(word[0]) == mark) && (s[i] == s[length-1]) {
				// Concatenate the current word with the previous word and remove the current word from the slice
				s[i-1] += word
				s = s[:length-1]
			}
		}
	}
	
	count := 0
	
	for i, char := range s {
		// If the character is "'", increment the counter
		if char == "'" && count == 0 {
			count += 1
			// Concatenate "'" with the next word and remove it from the current position
			s[i+1] = char + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	

	for i, char := range s {
		// If the character is "'", concatenate it with the previous word and remove it from the current position
		if char == "'" {
			s[i-1] = s[i-1] + char
			s = append(s[:i], s[i+1:]...)
		}
	}
	
	// Return the modified slice
	return s
}
