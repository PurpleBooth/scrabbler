package scrabbler

import (
	"testing"
)

func TestWord_string(t *testing.T) {
	word := MakeAlphabet().makeWord("cats")
	expected := "CATS"
	actual := word.String()

	if actual != expected {
		t.Error("Expected", expected, "got", actual)
	}

}
