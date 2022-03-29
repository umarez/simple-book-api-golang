package model

import "encoding/json"

type BookBody struct {
	Title    string      `json:"title" validate:"required"`
	Price    json.Number `json:"price" validate:"required,numeric"`
	SubTitle string      `json:"subtitle"`
}
