package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/betorvs/event-logger/appcontext"
	localtest "github.com/betorvs/event-logger/test"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestPostEvent(t *testing.T) {

	// Setup
	e1 := echo.New()
	req1 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(localtest.TestBadJSON))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	c1 := e1.NewContext(req1, rec1)
	c1.SetPath("/event-logger/v1/event")

	// Assertions
	if assert.NoError(t, PostEvent(c1)) {
		assert.Equal(t, http.StatusBadRequest, rec1.Code)
	}

	// with mock logger
	appcontext.Current.Add(appcontext.Logger, localtest.InitMockLogger)
	e2 := echo.New()
	req2 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(localtest.TestValidJSON))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	c2 := e2.NewContext(req2, rec2)
	c2.SetPath("/event-logger/v1/event")

	// Assertions
	if assert.NoError(t, PostEvent(c2)) {
		assert.Equal(t, http.StatusOK, rec2.Code)
	}

	e3 := echo.New()
	req3 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(localtest.TestInvalidJSON))
	req3.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec3 := httptest.NewRecorder()
	c3 := e3.NewContext(req3, rec3)
	c3.SetPath("/event-logger/v1/event")

	// Assertions
	if assert.NoError(t, PostEvent(c3)) {
		assert.Equal(t, http.StatusBadRequest, rec3.Code)
	}
}
