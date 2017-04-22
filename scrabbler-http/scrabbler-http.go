package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/newrelic/go-agent"
	"github.com/purplebooth/scrabbler"
	"log"
	"net/http"
	"os"
	"sort"
)

var dictionary scrabbler.Dictionary
var application newrelic.Application

func main() {
	nrName := os.Getenv("NEW_RELIC_APP_NAME")
	nrLicense := os.Getenv("NEW_RELIC_LICENSE_KEY")

	dictionary = scrabbler.MakeDictionary(scrabbler.MakeAlphabet())
	wordListFile := os.Args[1]

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

	(HttpServer{
		Port:      8080,
		NrName:    nrName,
		NrLicense: nrLicense,
	}).Listen()

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

type HttpServer struct {
	Port      int
	NrName    string
	NrLicense string
}

func (s HttpServer) Listen() {
	nrName := os.Getenv("NEW_RELIC_APP_NAME")
	nrLicense := os.Getenv("NEW_RELIC_LICENSE_KEY")

	router := httprouter.New()

	if nrName != "" && nrLicense != "" {
		config := newrelic.NewConfig(nrName, nrLicense)
		var err error
		application, err = newrelic.NewApplication(config)

		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

	}

	router.GET("/:word", s.wordHandler)
	router.GET("/", s.defaultPage)

	log.Println(fmt.Sprintf("Listening on :%v", s.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", s.Port), router))
}

type Response struct {
	Word  string
	Score int
}

func (s HttpServer) defaultPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if application != nil {
		txn := application.StartTransaction("/", w, r)
		defer txn.End()
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "max-age=2592000")

	fmt.Fprint(w, "Everything is okay! see /someword")
}

func (s HttpServer) wordHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if application != nil {
		txn := application.StartTransaction("/:word", w, r)
		defer txn.End()
	}

	log.Println("/:word")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "max-age=2592000")

	foundWords := dictionary.GetWords(ps.ByName("word"))
	sort.Sort(sort.Reverse(foundWords))

	response := make([]Response, 0)

	for _, foundWord := range foundWords {
		response = append(response, Response{Word: foundWord.Word.String(), Score: foundWord.GetScore()})
	}

	wordsJson, _ := json.Marshal(response)
	w.Write(wordsJson)
}
