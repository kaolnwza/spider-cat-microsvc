package main

import (
	handler "github.com/kaolnwza/spider-cat-microsvc/spider-service/handlers"
	repository "github.com/kaolnwza/spider-cat-microsvc/spider-service/repositories"
	service "github.com/kaolnwza/spider-cat-microsvc/spider-service/services"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	spiderRepo := repository.NewSpiderRepository()
	spiderSvc := service.NewSpiderService(spiderRepo)
	spiderHdr := handler.NewSpiderHandler(spiderSvc)

	e.GET("/", spiderHdr.GetSpidersHandler)

	e.Logger.Fatal(e.Start(":3333"))
}
