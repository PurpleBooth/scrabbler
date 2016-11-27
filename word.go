package scrabbler

import "strings"

type Word struct {
	letters []Letter
}

func (w Word) String() string {
	lettersArray := make([]string, 0)

	for _, letter := range w.letters {
		lettersArray = append(lettersArray, letter.letter)
	}

	return strings.Join(lettersArray, "")
}
