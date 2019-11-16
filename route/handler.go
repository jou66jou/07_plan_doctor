package route

import (
	"net/http"

	"github.com/labstack/echo"
)

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pin pin~~~poooooon!")
}
