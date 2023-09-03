package repository

import (
	"github.com/google/uuid"
	"github.com/kaolnwza/spider-cat-microsvc/inventory-service/model"
)

type inventoryRepo struct {
}

type InventoryRepository interface {
	IncreaseItem() error
	DecreaseItem() error
	GetInventory() ([]model.Inventory, error)
}

func NewInventoryRepository() inventoryRepo {
	return inventoryRepo{}
}

func (r inventoryRepo) IncreaseItem() error {
	return nil
}

func (r inventoryRepo) DecreaseItem() error {
	return nil
}

func (r inventoryRepo) GetInventory() ([]model.Inventory, error) {
	return []model.Inventory{
		{ItemUUID: uuid.New(), Amount: 1},
	}, nil
}
