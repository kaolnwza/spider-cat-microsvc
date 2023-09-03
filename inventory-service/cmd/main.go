package main

import (
	handler "github.com/kaolnwza/spider-cat-microsvc/inventory-service/handlers"
	repository "github.com/kaolnwza/spider-cat-microsvc/inventory-service/repositories"
	service "github.com/kaolnwza/spider-cat-microsvc/inventory-service/services"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	invRepo := repository.NewInventoryRepository()
	invSvc := service.NewInventoryService(invRepo)
	invHdr := handler.NewInventoryHandler(invSvc)

	e.GET("/", invHdr.GetInventoryHandler)

	e.Logger.Fatal(e.Start(":2222"))
}
