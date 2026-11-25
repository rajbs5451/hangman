package helpers

import (
	"fmt"
	"strings"
)

func IsLetterInString(letter, word string) bool {
	present := false
	w := strings.ToUpper(word)
	l := strings.ToUpper(letter)
	//return strings.Contains(w, l)

	for i, k := range w {
		if string(k) == l {
			present = true
			fmt.Println(i,k)
			break
		}

	}
	
	return present
}
