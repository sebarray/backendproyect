package deckdb

import "github.com/sebarray/backendproyect/model"

type DeckDB struct{}

type IDeck interface {
	CreateDeck(deck model.Deck) error
	DeleteDeck(id string) error
	UpdateDeck(deck model.Deck) error
	ReadDeck(def, id string) ([]model.Deck, error)
}

var providers = map[string]IDeck{}

func init() {
	providers["mysql"] = DeckDB{}
}

func GetProvider(typeDb string) IDeck {
	return providers[typeDb]
}
