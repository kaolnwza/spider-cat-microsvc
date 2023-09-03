package main

import (
	handler "github.com/kaolnwza/spider-cat-microsvc/cat-service/handlers"
	repository "github.com/kaolnwza/spider-cat-microsvc/cat-service/repositories"
	service "github.com/kaolnwza/spider-cat-microsvc/cat-service/services"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	healthHdr := handler.NewHealthHandler()

	catRepo := repository.NewCatRepository()
	catSvc := service.NewCatService(catRepo)
	catHdr := handler.NewCatHandler(catSvc)

	e.GET("/health", healthHdr.HealthCheck)
	e.GET("/", catHdr.GetCatHPHandler)

	e.Logger.Fatal(e.Start(":1111"))
}
