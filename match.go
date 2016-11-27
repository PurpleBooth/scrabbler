package scrabbler

import "strings"

type MatchList []Match

type TileModifier struct {
	multiplier int
}

type Match struct {
	Word
	tiles []TileModifier
}

func (a Alphabet) makeMatch(word Word, searchTerm string) Match {
	tileModifiers := make([]TileModifier, 0)
	match := Match{Word: word, tiles: tileModifiers}

	for _, letter := range word.letters {
		if strings.Contains(searchTerm, letter.letter) {
			match.tiles = append(match.tiles, TileModifier{multiplier: 1})
		} else {
			match.tiles = append(match.tiles, TileModifier{multiplier: 0})
		}
	}

	return match
}

func (m Match) GetScore() int {
	score := 0

	for index, letter := range m.letters {
		score = score + (letter.basePoints * m.tiles[index].multiplier)
	}

	return score
}

func (s MatchList) Len() int {
	return len(s)
}

func (s MatchList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s MatchList) Less(i, j int) bool {
	return s[i].GetScore() < s[j].GetScore()
}
