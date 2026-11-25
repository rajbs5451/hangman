package helpers

import (
	"fmt"
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
	for i, r := range test {
		A := IsLetterInString(r.letter, r.word)
		fmt.Println(i, r)
		if A != r.want {
			t.Error("ffddd")
		}
	}

}
