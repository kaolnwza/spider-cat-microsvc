package main

import (
	"github.com/kaolnwza/spider-cat-microsvc/cat-service/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	healthHdr := handler.NewHealthHandler()

	e.GET("/health", healthHdr.HealthCheck)

	e.Logger.Fatal(e.Start(":1111"))
}
