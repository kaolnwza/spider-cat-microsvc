package handler

import (
	"net/http"

	service "github.com/kaolnwza/spider-cat-microsvc/spider-service/services"
	"github.com/labstack/echo/v4"
)

type spiderHdr struct {
	spiderSvc service.SpiderService
}

func NewSpiderHandler(spiderSvc service.SpiderService) spiderHdr {
	return spiderHdr{spiderSvc: spiderSvc}
}

func (h spiderHdr) GetSpidersHandler(c echo.Context) error {
	spider, err := h.spiderSvc.GetSpidersService()
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.Logger().Print(spider)
	return c.JSON(http.StatusOK, spider)
}
