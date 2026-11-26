package helpers

import (
	"errors"
	"regexp"
	"strings"
)

var ErrLetterNotFound = errors.New("Letter not found in string")

func IsLetterInString(letter, word string) bool {
	// present := false
	w := strings.ToUpper(word)
	l := strings.ToUpper(letter)
	return strings.Contains(w, l)

	// for i, k := range w {
	// 	if string(k) == l {
	// 		present = true
	// 		fmt.Println(i,k)
	// 		break
	// 	}

	// }

	// return present
}

func IndexOfLetter(word, letter string) ([]int, error) {
	pos := make([]int, 0, len(word))
	word = strings.ToLower(word)
	letter = strings.ToLower(letter)
	for i, ch := range word {
		// fmt.Println(i, word, string(ch), letter)
		if letter == string(ch) {
			pos = append(pos, i)
		}
	}
	if len(pos) == 0 {
		return nil, ErrLetterNotFound
	}

	return pos, nil
}

func IndexOfLetterRegex(word, letter string) ([]int, error) {
	pos := make([]int, 0, len(word))
	re := regexp.MustCompile(`(?i)` + regexp.QuoteMeta(letter))

	for _, match := range re.FindAllStringIndex(word, -1) {
		pos = append(pos, match[0])
	}
	if len(pos) == 0 {
		return nil, ErrLetterNotFound
	}

	return pos, nil
}

func StringToDashes(w string) string {
	var dashes []string
	for range w {
		dashes = append(dashes, "_")
	}
	return strings.Join(dashes, " ")
}
