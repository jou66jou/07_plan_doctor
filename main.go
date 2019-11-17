package main

import (
	"agr-hack/internal/errors"
	"agr-hack/internal/format"
	"agr-hack/internal/rice"
	"agr-hack/route"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/labstack/echo"
)

type Fields map[string]interface{}

func main() {

	e := echo.New()
	// setting
	s := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
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

	// route set
	route.InitHandler(e)
	// data read
	err := rice.Init()
	if err != nil || rice.RiceInfo == nil {
		log.Println("rice data open fail : ", err)
		return
	}
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

func getIP() (string, error) {
	ip := ""
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("ipconfig")
		d, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
		ip = string(d)
	case "linux", "darwin":
		cmd := exec.Command("bash", "-c", "ifconfig | grep -Eo 'inet (addr:)?([0-9]*\\.){3}[0-9]*' | grep -Eo '([0-9]*\\.){3}[0-9]*' | grep -v '127.0.0.1'")
		d, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
		ip = string(d)
	}
	// ip has content
	if ip != "" {
		log.Println(ip)
		return ip, nil
	}
	return "", errors.New("your os is " + runtime.GOOS + " not support get ip command")
}
