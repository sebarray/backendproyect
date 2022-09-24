package model

type DeckCard struct {
	Id     string `json:"id"`
	IdDeck string `json:"iDeck"`
	IdUser string
	Cards  []Card
}
