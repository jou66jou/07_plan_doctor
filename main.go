package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
)

const (
	apiKey = "CWB-B598382E-A64D-4809-B598-5C434E4FCEAB"

	oneWeekWeather = "https://opendata.cwb.gov.tw/fileapi/v1/opendataapi/F-A0010-001?Authorization=CWB-B598382E-A64D-4809-B598-5C434E4FCEAB&downloadType=WEB&format=JSON"
)
type Fields map[string]interface{}

func main() {
	e := echo.New()
	// setting
	s := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
	}
	e.Debug = false 
	e.HTTPErrorHandler = errors.HTTPErrorHandlerForEcho
	// cover all api error response
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				logFields := Fields{}

				// get request data
				req := c.Request()
				{
					logFields["requestMethod"] = req.Method
					logFields["requestURL"] = req.URL.String()
				}

				str := fmt.Sprintf("%+v, error message : %+v\n", logFields, err)
				msg := format.GetCMDColor(format.Color_red, "[API ERROR] ")
				log.Printf(msg + str)
			}
			return err
		}
	})

	// route
	e.POST("/upload", upload)

	// test file
	dst, err := os.Create("test")
	if err != nil {
		log.Println("test file fail: " + err.Error())
	}
	dst.Close()

	// get local ip address by command
	address, err := getIP()
	if err != nil {
		log.Println("get ip fail : ", err)
	}
	if address == "" {
		log.Println("not found local ip address")
	}
	log.Println("local ip address :", format.GetCMDColor(format.Color_green, address))
	// start
	e.Logger.Fatal(e.StartServer(s))
}
