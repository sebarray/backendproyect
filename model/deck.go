package model

type Deck struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	IdUser string
}
