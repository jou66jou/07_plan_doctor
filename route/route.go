package route

import "github.com/labstack/echo"

func InitHandler(e *echo.Echo) {
	e.GET("/", ping)
	e.GET("/getWarning", getWarning)
	e.GET("/getHistory", getHistory)
}
