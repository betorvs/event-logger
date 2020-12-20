package controller

import (
	"net/http"

	"github.com/betorvs/event-logger/domain"
	"github.com/betorvs/event-logger/usecase"
	"github.com/labstack/echo/v4"
)

// PostEvent controller to receive an event
func PostEvent(c echo.Context) (err error) {
	event := new(domain.Event)
	if err = c.Bind(event); err != nil {
		return c.JSON(http.StatusBadRequest, "cannot load event data")
	}
	err = usecase.LogEvent(event)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "should contain name and/or kind")
	}

	return nil
}
