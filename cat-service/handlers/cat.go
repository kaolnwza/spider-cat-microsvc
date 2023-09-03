package handler

import (
	"net/http"

	service "github.com/kaolnwza/spider-cat-microsvc/cat-service/services"
	"github.com/labstack/echo/v4"
)

type catHdr struct {
	catSvc service.CatService
}

func NewCatHandler(catSvc service.CatService) catHdr {
	return catHdr{catSvc: catSvc}
}

func (h catHdr) GetCatHPHandler(c echo.Context) error {
	cat, err := h.catSvc.GetCatHPService()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Print(cat)
	return c.JSON(http.StatusOK, cat)
}
