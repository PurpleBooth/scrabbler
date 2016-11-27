package scrabbler

import "testing"

func TestMakeAlphabet(t *testing.T) {
	alphabet := MakeAlphabet()
	expected := map[string]int{
		"E": 1,
		"A": 1,
		"I": 1,
		"O": 1,
		"N": 1,
		"R": 1,
		"T": 1,
		"L": 1,
		"S": 1,
		"U": 1,
		"D": 2,
		"G": 2,
		"B": 3,
		"C": 3,
		"M": 3,
		"P": 3,
		"F": 4,
		"H": 4,
		"V": 4,
		"W": 4,
		"Y": 4,
		"K": 5,
		"J": 8,
		"X": 8,
		"Q": 10,
		"Z": 10,
	}

	for letter, point := range expected {
		if _, isSet := alphabet[letter]; !isSet {
			t.Error(
				"Not set", letter,
			)
		}

		if actualLetter := alphabet[letter]; actualLetter.basePoints != point {
			t.Error(
				"Expected points", point,
				"Got points", actualLetter.basePoints,
			)
		}
	}
}

func TestMakeWord(t *testing.T) {
	alphabet := MakeAlphabet()
	word := alphabet.makeWord("cat")
	expected := [3]Letter{
		{basePoints: 3, letter: "C"},
		{basePoints: 1, letter: "A"},
		{basePoints: 1, letter: "T"},
	}

	for index, expectedLetter := range expected {
		if actualLetter := word.letters[index]; actualLetter.letter != expectedLetter.letter {
			t.Error(
				"Expected letter", expectedLetter.letter,
				"Got letter", actualLetter.letter,
			)
		}

		if actualLetter := word.letters[index]; actualLetter.basePoints != expectedLetter.basePoints {
			t.Error(
				"Expected points", expectedLetter.basePoints,
				"Got points", actualLetter.basePoints,
			)
		}
	}
}
