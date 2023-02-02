package requests

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wkey-core/src/definition"
)

func (service *Service) Header(key, value string) *Service {
	if _, ok := service.headers[key]; ok {
		service.headers[key] = append(service.headers[key], value)
	} else {
		service.headers[key] = []string{value}
	}
	return service
}

func (service *Service) Headers(headers map[string][]string) *Service {
	service.headers = headers
	return service
}

func (service *Service) Cookie(key, value string) *Service {
	service.cookies[key] = value
	return service
}

func (service *Service) Cookies(cookies map[string]string) *Service {
	service.cookies = cookies
	return service
}

func (service *Service) BasicAuth(username, password string) *Service {
	service.username = username
	service.password = password
	service.isBasicAuth = true
	return service
}

func (service *Service) Send() ([]byte, error) {
	var bodyBuffer *bytes.Buffer
	if service.body != nil {
		if value, ok := service.body.([]byte); ok {
			bodyBuffer = bytes.NewBuffer(value)
		} else {
			bodyInBytes, err := json.Marshal(service.body)
			if err != nil {
				return nil, err
			}

			bodyBuffer = bytes.NewBuffer(bodyInBytes)
		}
	}

	ctx, cancel := service.ctx()
	defer cancel()

	var request *http.Request
	var err error
	if service.body != nil {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, bodyBuffer)
	} else {
		request, err = http.NewRequestWithContext(ctx, service.method, service.url, nil)
	}
	if err != nil {
		return nil, err
	}

	service.request = request

	service.fillHeaders()
	service.fillCookies()
	service.setBasicAuth()

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			definition.Logger.Error(err, "Close response body error")
		}
	}()

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	service.status = response.StatusCode

	return responseInBytes, nil
}

func (service *Service) SendStatus() ([]byte, int, error) {
	response, err := service.Send()
	return response, service.status, err
}
