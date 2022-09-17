package model

type Card struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Image       string `json:"image"`
	Description string `json:"description"`
	IdUser      string `json:"iduser"`
}
