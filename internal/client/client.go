package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

type Client interface {
	Sample() error
}

// Client 客製化 url.URL 和 http.Client
type client struct {
	BaseURL    *url.URL
	businessID string
	httpClient *http.Client
}

// NewClient 建立 Client
func NewClient(scheme, host, apiKey, businessID string, transport *http.Transport) Client {
	return &client{
		BaseURL: &url.URL{
			Scheme: scheme,
			Host:   host,
		},
		businessID: businessID,
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
