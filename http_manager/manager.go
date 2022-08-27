package http_manager

import (
	"encoding/base64"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

type httpManager struct {
	BaseURL       string
	Client        *fasthttp.Client
	ApiToken      *string
	ApiTokenKey   *string
	Username      *string
	Password      *string
	OkStatusCodes []int
}

func NewHttpManager(baseUrl string, accessToken, accessTokenKey, username, password *string, okStatusCodes ...int) *httpManager {
	return &httpManager{
		BaseURL: baseUrl,
		Client: &fasthttp.Client{
			MaxIdleConnDuration: 30 * time.Second,
		},
		ApiToken:      accessToken,
		ApiTokenKey:   accessTokenKey,
		Username:      username,
		Password:      password,
		OkStatusCodes: okStatusCodes,
	}
}

func (m *httpManager) sendRequest(method, url string, body []byte, req *fasthttp.Request) ([]byte, error) {
	url = m.BaseURL + url
	req.SetRequestURIBytes([]byte(url))

	fmt.Println(string(req.Header.Header()))

	res := fasthttp.AcquireResponse()
	if err := m.Client.Do(req, res); err != nil {
		log.Printf(`error when send request: method: %s, url: %s, body: %s, error: %s\n`, method, url, string(body), err.Error())
		return nil, newBadRequestError("error sending request: %s", err)

	}

	if !m.contains(res.StatusCode()) {
		log.Printf(`error when send request: method: %s, url: %s, body: %s, status_code: %v\n`, method, url, string(body), res.StatusCode())
		return nil, newBadRequestError("response or response body is null: response: %v, status_code: %v", res == nil, res.StatusCode())
	}

	if res == nil {
		return nil, newBadRequestError("response or response body is null: response: %v, body: %v", res == nil, res.Body)
	}

	fasthttp.ReleaseRequest(req)

	return res.Body(), nil
}

func (m *httpManager) contains(val int) bool {
	for _, i := range m.OkStatusCodes {
		if i == val {
			return true
		}
	}
	return false
}

func (m *httpManager) basicAuth(username, password string) string {
	auth := username + ":" + password
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(auth))
}
