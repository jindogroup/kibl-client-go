package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	KiblTimestampFormat = "2006-01-02T15:04:05.999Z"
	KiblRequestID       = "X-REQUEST-ID"
)

type Headers map[string]string

func (p *Headers) Set(request *http.Request) {
	if request == nil || len(*p) == 0 {
		return
	}
	for key, value := range *p {
		request.Header.Set(key, value)
	}
}

type Params map[string]string

func (p *Params) Add(key, value string) *Params {
	(*p)[key] = value
	return p
}
func (p *Params) AddIfMissing(key, value string) *Params {
	if _, ok := (*p)[key]; !ok {
		(*p)[key] = value
	}
	return p
}

func (p *Params) SetQueryParams(_url *url.URL) {
	if len(*p) == 0 {
		return
	}
	query := _url.Query()
	for k, v := range *p {
		query.Add(k, v)
	}
	_url.RawQuery = query.Encode()
}

func (p *Params) Request(method, path string) (request *http.Request, err error) {
	switch method {
	case http.MethodGet:
		_url, err := url.Parse(path)
		if err != nil {
			return nil, err
		}
		p.SetQueryParams(_url)
		return http.NewRequest(method, _url.String(), nil)
	case http.MethodPost:
		data, err := json.Marshal(p)
		if err != nil {
			return nil, err
		}
		return http.NewRequest(method, path, bytes.NewReader(data))
	default:
		panic("method " + method + " unsupported")
	}
}

func (p *Params) Merge(map[string]interface{}) *Params {

	return p
}
