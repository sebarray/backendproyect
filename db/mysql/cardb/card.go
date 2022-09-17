package cardb

import "github.com/sebarray/backendproyect/model"

type CardDB struct{}

type ICard interface {
	CreateCard(card model.Card) error
	DeleteCard(id string) error
	UpdateCard(card model.Card) error
	ReadCard(def, id string) ([]model.Card, error)
}

var providers = map[string]ICard{}

func init() {
	providers["mysql"] = CardDB{}
}

func GetProvider(typeDb string) ICard {
	return providers[typeDb]
}
