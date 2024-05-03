package reloaded

// check if a word is "a" or "A" and if the next word starts with a vowel.
// If these conditions are met, it replaces "a" with "an" or "A" with "An" accordingly.
func AddVowel(s []string) []string {
	for i, word := range s {
		if word == "a" && isVowel(s[i+1][0]) {
			s[i] = "an"
		} else if word == "A" && isVowel(s[i+1][0]) {
			s[i] = "An"
		}
	}
	return s
}

// It iterates over a list of known vowels, both lowercase and uppercase.
// If the character matches any of them, it returns true; otherwise, it returns false.
func isVowel(c byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}
	for _, vowel := range vowels {
		if c == vowel {
			return true
		}
	}
	return false
}
