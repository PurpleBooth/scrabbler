package scrabbler

import (
	"github.com/kokardy/listing"
	"sort"
	"strings"
)

type Dictionary struct {
	Alphabet Alphabet
	words    map[string][]Word
}

func MakeDictionary(Alphabet Alphabet) Dictionary {
	return Dictionary{Alphabet: Alphabet, words: map[string][]Word{}}
}

func (d Dictionary) AddWord(word string) {
	word = strings.TrimSpace(word)
	sortedWord := strings.ToUpper(alphabetiseWord(word))
	dictionaryWord := d.Alphabet.makeWord(word)

	d.words[sortedWord] = append(d.words[sortedWord], dictionaryWord)
}

func (d Dictionary) GetWords(searchWord string) MatchList {
	searchWord = strings.ToUpper(searchWord)
	matchingWords := make(MatchList, 0)
	searchWords := d.getSearchWords(searchWord)

	for _, wordCombination := range searchWords {

		if foundWords, hasFoundWords := d.words[wordCombination]; hasFoundWords {
			for _, foundWord := range foundWords {
				matchingWords = append(matchingWords, d.Alphabet.makeMatch(foundWord, searchWord))
			}
		}
	}

	return matchingWords
}

func alphabetiseWord(word string) string {
	splitWord := strings.Split(word, "")
	sort.Strings(splitWord)
	sortedWord := strings.Join(splitWord, "")

	return sortedWord
}

func (d Dictionary) getSearchWords(searchWord string) []string {
	searchWords := make([]string, 0)

	if strings.Contains(searchWord, " ") {
		blankTileCount := strings.Count(searchWord, " ")
		wordWithoutSpaces := strings.Replace(searchWord, " ", "", -1)
		blankTilesRepresent := make([]string, 0)

		for i := 0; i < blankTileCount; i++ {
			for letter := range d.Alphabet {
				blankTilesRepresent = append(blankTilesRepresent, letter)
			}
		}

		stringCombiner := make(listing.StringReplacer, 0)
		stringCombiner = append(stringCombiner, blankTilesRepresent...)

		for blankTileCombinations := range listing.Combinations(stringCombiner, blankTileCount, false, blankTileCount) {
			wordCombination := make([]string, 0)

			wordCombination = append(wordCombination, blankTileCombinations.(listing.StringReplacer)...)
			wordCombination = append(wordCombination, wordWithoutSpaces)
			searchWords = append(searchWords, strings.ToUpper(alphabetiseWord(strings.Join(wordCombination, ""))))
		}
	} else {
		searchWords = append(searchWords, strings.ToUpper(searchWord))
	}

	wordCombinations := map[string]bool{}

	for _, searchWord := range searchWords {
		stringSorter := listing.StringReplacer(strings.Split(alphabetiseWord(searchWord), ""))

		for combinationSize := 1; combinationSize <= len(searchWord); combinationSize++ {
			for combinations := range listing.Combinations(stringSorter, combinationSize, false, len(searchWord)) {
				combination := strings.Join(combinations.(listing.StringReplacer), "")

				wordCombinations[combination] = true
			}
		}
	}

	matchingWords := make([]string, 0)

	for wordCombination, _ := range wordCombinations {
		matchingWords = append(matchingWords, wordCombination)
	}

	sort.Strings(matchingWords)

	return matchingWords
}
