package handler

import (
	"net/http"

	"github.com/google/uuid"
	service "github.com/kaolnwza/spider-cat-microsvc/cat-service/services"
	"github.com/labstack/echo/v4"
)

type spiderHandler struct {
	spiderSvc service.SpiderService
}

func NewSpiderHandler(spiderSvc service.SpiderService) spiderHandler {
	return spiderHandler{spiderSvc: spiderSvc}
}

func (h spiderHandler) AttackSpiderHandler(c echo.Context) error {
	spiderUUID, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := h.spiderSvc.AttackSpiderService(spiderUUID); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"msg": "เอื้อออ"})
}
