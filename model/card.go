package model

import "github.com/google/uuid"

type Card struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
}
