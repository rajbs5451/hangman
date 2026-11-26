package helpers

import (
	"fmt"
	"slices"
	"testing"
)

// func TestIsLetterInString(t *testing.T){
// 	word := "foot"
// 	letter := "t"
// 	got := IsLetterInString(letter, word)
// 	want := true

// 	if got != want{
// 		t.Errorf("Got %v, want %v\n",got,want)
// 	}
// }

func TestIsLetterInString(t *testing.T) {
	test := []struct {
		word, letter string
		want         bool
	}{
		{"raj", "f", false},
		{"Sun", "U", true},
		{"moon", "M", true},
		{"Rat", "R", true},
	}
	for _, tc := range test {
		got := IsLetterInString(tc.letter, tc.word)
		if got != tc.want {
			t.Errorf("Word: %q, letter: %q, expected: %v, got=%v\n", tc.word, tc.letter, tc.want, got)
		}
	}

}

func TestIndexOfLetter(t *testing.T) {
	tests := []struct {
		word, letter string
		want         []int
		err          error
	}{
		{"raj", "f", nil, ErrLetterNotFound},
		{"Sun", "U", []int{1}, nil},
		{"moon", "O", []int{1, 2}, nil},
	}
	for _, tc := range tests {
		got, err := IndexOfLetterRegex(tc.word, tc.letter)
		if err == nil && tc.err == ErrLetterNotFound {
			t.Errorf("Word: %q, letter: %q, expected: error: %v, got=%v\n", tc.word, tc.letter, tc.err.Error(), err)
		} else if !slices.Equal(got, tc.want) {
			t.Errorf("Word: %q, letter: %q, expected: %v, got=%v\n", tc.word, tc.letter, tc.want, got)
		}
	}

}

func ExampleIndexOfLetter() {
	pos, err := IndexOfLetter("rajendra", "f")
	fmt.Println(pos, err.Error())
	// Output: [] Letter not found in string

}
func ExampleStringToDashes() {
	input := "rajendra"
	// want := "_ _ _ _ _ _ _ _"
	got := StringToDashes(input)
	fmt.Println(got)
	// Output: _ _ _ _ _ _ _ _
}
