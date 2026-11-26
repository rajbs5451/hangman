package main

import (
	"fmt"
	"os"

	"github.com/rajbs5451/hangman/helpers"
)

func main() {
	word := os.Args[1]
	out := helpers.StringToDashes(word)
	fmt.Println(out)
}
