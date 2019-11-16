package route

import (
	"agr-hack/internal/client"
	"agr-hack/internal/errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/labstack/echo"
)

const (
	apiKey = "CWB-B598382E-A64D-4809-B598-5C434E4FCEAB"

	apiURL = "https://opendata.cwb.gov.tw/"
)

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pin pin~~~poooooon!")
}

func getOneWeekWeather(c echo.Context) error {
	u, _ := url.Parse(apiURL)

	client := client.NewClient(
		u.Scheme,
		u.Host,
		&http.Transport{
			MaxConnsPerHost:     24,
			MaxIdleConnsPerHost: 24,
			MaxIdleConns:        48,
			IdleConnTimeout:     60 * time.Second,
		},
	)
	data, resp, err := client.OneWeekWeather()
	if err != nil {
		s := fmt.Sprintf("fail request url : %s , err : %+v", resp.Request.URL.String(), err)
		return errors.NewWithMessage(errors.ErrInternalError, s)
	}
	return c.JSON(http.StatusOK, data)
}
