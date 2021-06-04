package server

import (
	"net/http"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/labstack/echo/v4"
)

type CustomHTTPErrorHandler struct {
	debug  bool
	logger echo.Logger
}

func (che CustomHTTPErrorHandler) Handler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
	}

	if ev, ok := err.(entity.ErrValidation); ok {
		he.Code = http.StatusBadRequest
		if ev.Message == "" {
			he.Message = http.StatusText(http.StatusBadRequest)
		} else {
			he.Message = ev.Message
		}
	} else if ent, ok := err.(entity.ErrNotFound); ok {
		he.Code = http.StatusNotFound
		if ev.Message == "" {
			he.Message = http.StatusText(http.StatusNotFound)
		} else {
			he.Message = ent.Message
		}
	}

	code := he.Code
	payload := he.Message
	if m, ok := he.Message.(string); ok {
		if che.debug {
			payload = echo.Map{"message": m, "error": err.Error()}
		} else {
			payload = echo.Map{"message": m}
		}
	}

	// Send response
	if !c.Response().Committed {
		if c.Request().Method == http.MethodHead {
			err = c.NoContent(he.Code)
		} else {
			err = c.JSON(code, payload)
		}
		if err != nil {
			che.logger.Error(err)
		}
	}
}
