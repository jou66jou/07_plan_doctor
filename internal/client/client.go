package client

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

const (
	apiKey = "CWB-B598382E-A64D-4809-B598-5C434E4FCEAB"
)

type Client interface {
	Sample() error
	OneWeekWeather() (*OneWeekWeatherResp, *http.Response, error)
}

// Client 客製化 url.URL 和 http.Client
type client struct {
	BaseURL    *url.URL
	businessID string
	httpClient *http.Client
}

// NewClient 建立 Client
func NewClient(scheme, host string, transport *http.Transport) Client {
	return &client{
		BaseURL: &url.URL{
			Scheme: scheme,
			Host:   host,
		},
		httpClient: &http.Client{
			Transport: transport,
		},
	}
}
func (c *client) doRequest(req *http.Request) (int, []byte, *http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, nil, &http.Response{}, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, resp, errors.Wrap(err, "read response body failed")
	}
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return 0, nil, resp, errors.New("get err")
	}
	return resp.StatusCode, b, resp, nil
}

func (c *client) Sample() error {
	return nil
}

type OneWeekWeatherResp struct {
	Data Locations `json:"records"`
}
type Locations struct {
	LocatEle []LocatEle `json:"locations"`
}
type LocatEle struct {
	LocatContent []Location `json:"location"`
}
type Location struct {
	Name     string    `json:"locationName"`   // 地區名
	LAT      string    `json:"lat"`            // 緯度
	LON      string    `json:"lon"`            // 經度
	Weathers []Weather `json:"weatherElement"` // 氣象資訊
}
type Weather struct {
	Desc  string `json:"description"`
	Times []Time `json:"time"`
}
type Time struct {
	Elements []Ele  `json:"elementValue"`
	Start    string `json:"startTime"`
	End      string `json:"endTime"`
}
type Ele struct {
	Mersures string `json:"measures"`
	Value    string `json:"value"`
}

func (c *client) OneWeekWeather() (*OneWeekWeatherResp, *http.Response, error) {
	var data = &OneWeekWeatherResp{}
	// 複製一份
	u := c.BaseURL.ResolveReference(&url.URL{Path: "api/v1/rest/datastore/F-D0047-091"})
	// 建立 request 並設定 Query String
	req, err := http.NewRequest(http.MethodGet, u.String(), http.NoBody)
	if err != nil {
		return data, &http.Response{}, errors.Wrap(err, "http NewRequest failed")
	}
	q := req.URL.Query()
	q.Add("Authorization", apiKey)
	req.URL.RawQuery = q.Encode()
	log.Printf("request url : %+v\n", req.URL.String())

	// 執行 HTTP Request
	_, b, resp, reqErr := c.doRequest(req)
	err = json.Unmarshal(b, data)
	if err != nil {
		return data, resp, errors.Wrap(err, "json Unmarshal failed : "+err.Error())
	}
	if reqErr != nil {
		return data, resp, errors.Wrap(reqErr, "c.doRequest failed"+reqErr.Error())
	}
	// 處理 Status 200 多的 response
	return data, resp, nil

}
