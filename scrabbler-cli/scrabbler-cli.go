package main

import (
	"bufio"
	"fmt"
	"github.com/purplebooth/scrabbler"
	"log"
	"os"
	"sort"
)

func main() {
	dictionary := scrabbler.MakeDictionary(scrabbler.MakeAlphabet())
	wordListFile := os.Args[1]
	tiles := os.Args[2]

	fmt.Println("Loading dictionary:", wordListFile)

	file, err := os.Open(wordListFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		dictionary.AddWord(word)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Searching words")
	words := dictionary.GetWords(tiles)
	sort.Sort(sort.Reverse(words))

	fmt.Println("Word", "\t", "Points")

	for _, match := range words {
		fmt.Println(match.String(), "\t", match.GetScore())
	}

	fmt.Println("Done!")
}
