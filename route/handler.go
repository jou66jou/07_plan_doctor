package route

import (
	"agr-hack/internal/client"
	"agr-hack/internal/errors"
	"agr-hack/internal/rice"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
)

const (
	apiKey = "CWB-B598382E-A64D-4809-B598-5C434E4FCEAB"
	apiURL = "https://opendata.cwb.gov.tw/"

	avgOccurDay = 3
)

func ping(c echo.Context) error {
	return c.JSON(http.StatusOK, "pin pin~~~poooooon!")
}

type WarningResp struct {
	Data []Warn `json:"data"`
}
type Warn struct {
	Disease       string         `json:"disease"` // 病害名稱
	Temp          string         `json:"temp"`    // 溫度
	RH            string         `json:"RH"`      // 濕度
	Comment       string         `json:"comment"` // 備註
	LocationInfos []LocationInfo `json:"locationInfos"`
}
type LocationInfo struct {
	LocationName string `json:"locationName"` // 地區名
	LAT          string `json:"lat"`          // 緯度
	LON          string `json:"lon"`          // 經度
}

type conditionCheck struct {
	rangeTemp *rangeTemp
	overTemp  *overTemp
	rh        *rh
	dry       *dry
}
type rangeTemp struct {
	low   int
	hight int
}
type overTemp struct {
	over int
}
type rh struct {
	rh int
}
type dry struct {
	rh int
}

func getWarning(c echo.Context) error {
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
	if data == nil {
		s := fmt.Sprintf("response data is nil, url : %s ", resp.Request.URL.String())
		return errors.NewWithMessage(errors.ErrInternalError, s)
	}
	result, err := warningAnalysis(data)
	if err != nil {
		return errors.NewWithMessage(errors.ErrInternalError, err.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func warningAnalysis(weathers *client.OneWeekWeatherResp) (resp *WarningResp, err error) {
	// var check bool
	for _, r := range rice.RiceInfo.Data {
		var warn = Warn{}
		rangeMonth := strings.Split(r.Occur, ",")
		// 若疾病未在現在月份則略過
		if !checkRangeMonth(rangeMonth) {
			continue
		}
		warn.Disease = r.Name
		warn.Temp = r.Temp
		warn.RH = r.RH
		warn.Comment = r.Comment
		// 判斷發病條件
		conditions := strings.Split(r.Condition, ",")
		local, err := checkConditions(r, conditions, weathers)
		if local != nil && err != nil {
			return nil, errors.NewWithMessage(errors.ErrInternalError, err.Error())
		}
		if warn.LocationInfos != nil {
			resp.Data = append(resp.Data, warn)
		}
	}

	return getResult()
}

func checkConditions(r rice.Data, conditions []string, weathers *client.OneWeekWeatherResp) (local *LocationInfo, err error) {
	// var locals []LocationInfo
	var conditionCheck conditionCheck
	for _, con := range conditions {
		switch con {
		case "temp": // 溫度之間
			if r.Temp != "" {
				var low, hight int
				temp := strings.Split(r.Temp, "-")
				if len(temp) > 1 {
					hight, err = strconv.Atoi(temp[1])
					if temp[0] != "" {
						low, err = strconv.Atoi(temp[0])
					} else {
						// 低溫無下限
						low = -100
					}
				} else {
					low, err = strconv.Atoi(temp[1])
					// 高溫無上限
					hight = 100
				}
				if err != nil {
					return nil, err
				}
				t := &rangeTemp{low: low, hight: hight}
				conditionCheck.rangeTemp = t
			}
		case "overTemp": // 溫度之差
			if r.Temp != "" {
				var low, hight int
				temp := strings.Split(r.Temp, "-")
				if len(temp) > 1 {
					hight, err = strconv.Atoi(temp[1])
					if temp[0] != "" {
						low, err = strconv.Atoi(temp[0])
					} else {
						// 低溫無下限
						low = -100
					}
				} else {
					low, err = strconv.Atoi(temp[1])
					// 高溫無上限
					hight = 100
				}
				if err != nil {
					return nil, err
				}
				t := &overTemp{over: hight - low}
				conditionCheck.overTemp = t
			}
		case "rh": // 濕度 >= 90
			rh := &rh{rh: 90}
			conditionCheck.rh = rh
		case "dry": // 濕度 < 80
			dry := &dry{rh: 90}
			conditionCheck.dry = dry
		}
	}
	// if conditionCheck.rangeTemp != nil {
	// 	for _, local := range weathers.Data.LocatEle {

	// 		if !getCheckTemp(conditionCheck.rangeTemp.low,conditionCheck.rangeTemp.hight,){
	// 			return
	// 		}
	// 	}
	// }
	// if conditionCheck.overTemp != nil {

	// }
	// if conditionCheck.rh != nil {

	// }
	// if conditionCheck.dry != nil {

	// }

	return
}

func getCheckTemp(low, hight int, tempEles *client.LocatEle) (combo bool) {
	return
}

func getOverTemp(tgr int, tempEles *client.LocatEle) (combo bool, avg int) {
	return true, 0
}

func checkRangeMonth(rangeMonth []string) bool {
	currentTime := time.Now()
	for _, r := range rangeMonth {
		months := strings.Split(r, "-")
		min, _ := strconv.Atoi(months[0])
		max, _ := strconv.Atoi(months[1])
		if min > int(currentTime.Month()) || max < int(currentTime.Month()) {
			continue
		}
		return true
	}
	return false
}

func getResult() (resp *WarningResp, err error) {
	resp = &WarningResp{}
	// Open our jsonFile
	jsonFile, err := os.Open("internal/rice/rice_result.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully Opened rice_result.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
