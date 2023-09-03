package handler

import (
	"net/http"

	service "github.com/kaolnwza/spider-cat-microsvc/inventory-service/services"
	"github.com/labstack/echo/v4"
)

type inventoryHdr struct {
	invSvc service.InventoryService
}

func NewInventoryHandler(invSvc service.InventoryService) inventoryHdr {
	return inventoryHdr{invSvc: invSvc}
}

func (h inventoryHdr) GetInventoryHandler(c echo.Context) error {
	inv, err := h.invSvc.GetInventoryService()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Print(inv)
	return c.JSON(http.StatusOK, inv)
}
