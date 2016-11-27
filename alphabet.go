package scrabbler

import (
	"strings"
)

type Letter struct {
	letter     string
	basePoints int
}

type Alphabet map[string]Letter

func MakeAlphabet() Alphabet {
	alphabet := map[string]Letter{}
	for point, letters := range map[int][]string{
		1:  {"E", "A", "I", "O", "N", "R", "T", "L", "S", "U"},
		2:  {"D", "G"},
		3:  {"B", "C", "M", "P"},
		4:  {"F", "H", "V", "W", "Y"},
		5:  {"K"},
		8:  {"J", "X"},
		10: {"Q", "Z"},
	} {
		for _, letter := range letters {
			alphabet[letter] = Letter{
				letter:     letter,
				basePoints: point,
			}
		}
	}

	return alphabet
}

func (a Alphabet) makeWord(word string) Word {
	letters := make([]Letter, 0)

	for _, element := range strings.Split(word, "") {
		letters = append(letters, a[strings.ToUpper(element)])
	}

	return Word{letters: letters}
}
