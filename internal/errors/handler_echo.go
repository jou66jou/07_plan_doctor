package errors

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// HTTPErrorHandlerForEcho responds error response according to given error.
func HTTPErrorHandlerForEcho(err error, c echo.Context) {
	if err == nil {
		return
	}

	_, ok := err.(*echo.HTTPError)
	if ok {
		_ = c.JSON(http.StatusInternalServerError, ErrInternalError)
		return
	}

	causeErr := errors.Cause(err)
	_err, ok := causeErr.(_error)
	errr := _error{}
	if !ok || _err == errr {
		_ = c.JSON(http.StatusInternalServerError, ErrInternalError)
		return
	}
	if len(_err.Code) < 3 {
		_ = c.JSON(http.StatusInternalServerError, ErrInternalError)
		return
	}
	_ = c.JSON(_err.Status, _err)
}

// NotFoundHandlerForEcho responds not found response.
func NotFoundHandlerForEcho(c echo.Context) error {
	return c.JSON(http.StatusNotFound, ErrResourceNotFound)
}
