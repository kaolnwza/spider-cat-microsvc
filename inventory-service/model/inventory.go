package model

import "github.com/google/uuid"

type Inventory struct {
	ItemUUID uuid.UUID `json:"item_uuid"`
	Amount   int       `json:"amount"`
}
