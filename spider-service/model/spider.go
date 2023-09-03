package model

import "github.com/google/uuid"

type Spider struct {
	UUID           uuid.UUID  `json:"uuid"`
	Name           string     `json:"name"`
	Damage         int        `json:"damage"`
	HP             int        `json:"hp"`
	SpiderAmount   int        `json:"spider_amount"`
	DropItemUUID   *uuid.UUID `json:"drop_item_uuid"`
	DropItemAmount *int       `json:"drop_item_amount"`
}
