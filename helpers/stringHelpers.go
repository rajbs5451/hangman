package helpers

import "strings"

func IsLetterInString(letter, word string) bool {
	// present := false
	w := strings.ToUpper(word)
	l := strings.ToUpper(letter)
	return strings.Contains(w, l)

	// for _, k := range w {
	// 	if string(k) == l {
	// 		present = true
	// 		break
	// 	}

	// }
	// return present
}
