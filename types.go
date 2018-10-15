package main

type Set struct {
	Name               string `json:"name"`
	Description        string `json:"description"`
	TermLanguage       string `json:"termLanguage"`
	DefinitionLanguage string `json:"definitionLanguage"`
	Cards              []Card `json:"cards"`
}

type Card struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
}
