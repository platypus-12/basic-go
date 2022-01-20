package word

import "unicode"

// func IsPalindrome(s string) bool {
// 	for i := range s {
// 		if s[i] != s[len(s)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
// ↑ バイト列を使ってしまっている

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
