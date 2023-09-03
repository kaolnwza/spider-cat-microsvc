package service

import (
	"github.com/kaolnwza/spider-cat-microsvc/inventory-service/model"
	repository "github.com/kaolnwza/spider-cat-microsvc/inventory-service/repositories"
)

type inventorySvc struct {
	invRepo repository.InventoryRepository
}

type InventoryService interface {
	IncreaseInventoryService() error
	GetInventoryService() ([]model.Inventory, error)
}

func NewInventoryService(invRepo repository.InventoryRepository) inventorySvc {
	return inventorySvc{
		invRepo: invRepo,
	}
}

func (s inventorySvc) IncreaseInventoryService() error {
	return nil
}

func (s inventorySvc) GetInventoryService() ([]model.Inventory, error) {
	inv, err := s.invRepo.GetInventory()
	if err != nil {
		return nil, err
	}

	return inv, nil
}
