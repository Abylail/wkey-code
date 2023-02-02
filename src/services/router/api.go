package router

import (
	"encoding/json"
	"github.com/lowl11/lazy-gateway/requests"
	"math/rand"
	"wkey-core/src/definition"
)

func Send(config *Config) ([]byte, error) {
	hosts := definition.Config.Hosts

	// read request body is exist
	var requestBodyInBytes []byte
	if config.Body != nil {
		var err error
		requestBodyInBytes, err = json.Marshal(config.Body)
		if err != nil {
			return nil, err
		}
	}

	// headers
	headers := make(map[string][]string)
	if config.Headers != nil {
		for key, value := range config.Headers {
			headers[key] = value
		}
	}

	// cookies
	cookies := make(map[string]string)
	if config.Cookies != nil {
		for key, value := range config.Cookies {
			cookies[key] = value
		}
	}

	// choose need host
	sendHost := hosts[rand.Intn(len(hosts))]
	sendUrl := sendHost + config.Port + config.Endpoint

	// send request
	response, status, err := requests.New(config.Method, sendUrl, requestBodyInBytes).
		Headers(headers).
		Cookies(cookies).
		SendStatus()
	if err != nil {
		return nil, err
	}

	_ = status

	return response, nil
}
