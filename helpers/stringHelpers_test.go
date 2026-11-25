package helpers

import "testing"

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
	tests := []struct {
		word, letter string
		want         bool
	}{
		{"foot", "o", true},
		{"foot", "z", false},
		{"foot", "O", true},
		{"Foot", "f", true},
		{"hands", "S", true},
	}

	for _, tc := range tests {
		got := IsLetterInString(tc.letter, tc.word)
		if got != tc.want {
			t.Errorf("Input: %q, Letter: %q \n Got %v, want %v\n", tc.word, tc.letter, got, tc.want)
		}

	}
}
