package main

import (
	"github.com/betorvs/event-logger/config"
	"github.com/betorvs/event-logger/controller"
	_ "github.com/betorvs/event-logger/gateway/customlog"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
