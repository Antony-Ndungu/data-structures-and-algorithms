package algorithms

import "strings"

// IsPalindrome checks if the string s is a Palindrome or not
func IsPalindrome(s string) bool {
	// Convert string s to lowercase
	s = strings.ToLower(s)
	// Remove spaces
	s = strings.ReplaceAll(s, " ", "")
	// Convert the string to a slice of runes
	r := []rune(s)
	// reverse the slice of runes
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		if r[i] != r[j] {
			return false
		}
	}
	return true
}

// 0(n) time | 0(1) space
