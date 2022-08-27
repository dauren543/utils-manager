package http_manager

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

func (m *httpManager) SendRequestBasicAuth(method, url string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes([]byte(method))
	req.Header.SetContentType("application/json")

	if m.Username != nil || m.Password != nil {
		req.Header.Set("Authorization", m.basicAuth(*m.Username, *m.Password))
	} else {
		return nil, errors.New("username or password is empty")
	}

	req.SetBodyRaw(body)

	return m.sendRequest(method, url, body, req)
}

func (m *httpManager) SendRequestOpen(method, url string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes([]byte(method))
	req.Header.SetContentType("application/json")

	req.SetBodyRaw(body)

	return m.sendRequest(method, url, body, req)
}

func (m *httpManager) SendRequestWithApiToken(method, url string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes([]byte(method))
	req.Header.SetContentType("application/json")

	if m.ApiTokenKey != nil || m.ApiToken != nil {
		req.Header.Set(*m.ApiTokenKey, *m.ApiToken)
	} else {
		return nil, errors.New("api token or api token key is empty")
	}

	req.SetBodyRaw(body)

	return m.sendRequest(method, url, body, req)
}

func (m *httpManager) SendRequestWithBearerToken(method, url, bearer string, body []byte) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethodBytes([]byte(method))
	req.Header.SetContentType("application/json")

	if bearer != "" {
		req.Header.Set("Authorization", fmt.Sprintf(`Bearer %v`, bearer))
	}

	req.SetBodyRaw(body)

	return m.sendRequest(method, url, body, req)
}
