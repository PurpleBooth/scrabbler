package scrabbler

import (
	"testing"
)

func TestMatch_getScore(t *testing.T) {
	alphabet := MakeAlphabet()
	word := alphabet.makeWord("CATS")
	match := alphabet.makeMatch(word, "C TS")
	expected := 5
	actual := match.GetScore()

	if actual != expected {
		t.Error(
			"Score invalid, expected",
			expected,
			"got",
			actual,
		)
	}
}
