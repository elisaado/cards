package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func parseFile() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}
}

func Shuffle(vals []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([]Card, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		ret[i] = vals[randIndex]
	}
	return ret
}

func remove(s []Card, c Card) []Card {
	var i int

	for j, card := range s {
		if card == c {
			i = j
		}
	}

	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
