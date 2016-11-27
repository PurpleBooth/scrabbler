package scrabbler

import (
	"testing"
)

func TestDictionary_AnagramSolver(t *testing.T) {
	dictionary := Dictionary{Alphabet: MakeAlphabet(), words: map[string][]Word{}}
	dictionary.AddWord("cat")
	dictionary.AddWord("cats")
	dictionary.AddWord("tac")
	dictionary.AddWord("tacs")
	dictionary.AddWord("bat")
	dictionary.AddWord("orange")
	dictionary.AddWord("bats")

	actual := dictionary.GetWords("tacs")
	expected := make([]Word, 0)
	expected = append(expected, dictionary.Alphabet.makeWord("cat"))
	expected = append(expected, dictionary.Alphabet.makeWord("cats"))
	expected = append(expected, dictionary.Alphabet.makeWord("tac"))
	expected = append(expected, dictionary.Alphabet.makeWord("tacs"))

	if len(actual) != len(expected) {
		t.Error(
			"Wrong number of words returned, expected",
			len(expected),
			expected,
			"got",
			len(actual),
			actual,
		)
	}
}
func TestDictionary_ScrabbleSolver(t *testing.T) {
	dictionary := Dictionary{Alphabet: MakeAlphabet(), words: map[string][]Word{}}
	dictionary.AddWord("cat")
	dictionary.AddWord("cats")
	dictionary.AddWord("tac")
	dictionary.AddWord("tacs")
	dictionary.AddWord("bat")
	dictionary.AddWord("orange")
	dictionary.AddWord("bats")

	actual := dictionary.GetWords("t cs")
	expected := make([]Word, 0)
	expected = append(expected, dictionary.Alphabet.makeWord("cat"))
	expected = append(expected, dictionary.Alphabet.makeWord("cats"))
	expected = append(expected, dictionary.Alphabet.makeWord("tac"))
	expected = append(expected, dictionary.Alphabet.makeWord("tacs"))

	if len(actual) != len(expected) {
		t.Error(
			"Wrong number of words returned, expected",
			len(expected),
			expected,
			"got",
			len(actual),
			actual,
		)
	}
}

func TestDictionarySearchWords(t *testing.T) {
	dictionary := Dictionary{Alphabet: MakeAlphabet(), words: map[string][]Word{}}
	actual := dictionary.getSearchWords("d g")
	expected := []string{"A", "AD", "ADG", "AG", "B", "BD", "BDG", "BG", "C", "CD", "CDG", "CG", "D", "DD", "DDG", "DE", "DEG", "DF", "DFG", "DG", "DGG", "DGH", "DGI", "DGJ", "DGK", "DGL", "DGM", "DGN", "DGO", "DGP", "DGQ", "DGR", "DGS", "DGT", "DGU", "DGV", "DGW", "DGX", "DGY", "DGZ", "DH", "DI", "DJ", "DK", "DL", "DM", "DN", "DO", "DP", "DQ", "DR", "DS", "DT", "DU", "DV", "DW", "DX", "DY", "DZ", "E", "EG", "F", "FG", "G", "GG", "GH", "GI", "GJ", "GK", "GL", "GM", "GN", "GO", "GP", "GQ", "GR", "GS", "GT", "GU", "GV", "GW", "GX", "GY", "GZ", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	if len(actual) != len(expected) {
		t.Error(
			"Wrong number of words returned, expected",
			len(expected),
			expected,
			"got",
			len(actual),
			actual,
		)
	}

	for i, value := range expected {
		if actual[i] != value {
			t.Error(
				"Word was wrong, expected",
				value,
				"got",
				actual[i],
			)
		}
	}
}

func BenchmarkDictionary_withSpace(b *testing.B) {
	dictionary := Dictionary{Alphabet: MakeAlphabet(), words: map[string][]Word{}}
	dictionary.AddWord("cat")
	dictionary.AddWord("cats")
	dictionary.AddWord("tac")
	dictionary.AddWord("tacs")
	dictionary.AddWord("bat")
	dictionary.AddWord("orange")
	dictionary.AddWord("bats")

	expected := make([]Word, 0)
	expected = append(expected, dictionary.Alphabet.makeWord("cat"))
	expected = append(expected, dictionary.Alphabet.makeWord("cats"))
	expected = append(expected, dictionary.Alphabet.makeWord("tac"))
	expected = append(expected, dictionary.Alphabet.makeWord("tacs"))

	b.ResetTimer()
	actual := dictionary.GetWords("T cs")

	if len(actual) != len(expected) {
		b.Error(
			"Wrong number of words returned, expected",
			len(expected),
			expected,
			"got",
			len(actual),
			actual,
		)
	}
}

func BenchmarkDictionary_withoutSpace(b *testing.B) {
	dictionary := Dictionary{Alphabet: MakeAlphabet(), words: map[string][]Word{}}
	dictionary.AddWord("cat")
	dictionary.AddWord("cats")
	dictionary.AddWord("tac")
	dictionary.AddWord("tacs")
	dictionary.AddWord("bat")
	dictionary.AddWord("orange")
	dictionary.AddWord("bats")

	expected := make([]Word, 0)
	expected = append(expected, dictionary.Alphabet.makeWord("cat"))
	expected = append(expected, dictionary.Alphabet.makeWord("cats"))
	expected = append(expected, dictionary.Alphabet.makeWord("tac"))
	expected = append(expected, dictionary.Alphabet.makeWord("tacs"))

	b.ResetTimer()
	actual := dictionary.GetWords("TACs")

	if len(actual) != len(expected) {
		b.Error(
			"Wrong number of words returned, expected",
			len(expected),
			expected,
			"got",
			len(actual),
			actual,
		)
	}
}
