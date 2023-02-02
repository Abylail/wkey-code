package requests

import (
	"context"
	"net/http"
	"time"
)

func (service *Service) fillHeaders() {
	if service.headers == nil {
		return
	}

	for key, value := range service.headers {
		for _, item := range value {
			service.request.Header.Add(key, item)
		}
	}
}

func (service *Service) fillCookies() {
	if service.cookies == nil {
		return
	}

	for key, value := range service.cookies {
		service.request.AddCookie(&http.Cookie{
			Name:  key,
			Value: value,
		})
	}
}

func (service *Service) setBasicAuth() {
	if !service.isBasicAuth {
		return
	}

	service.request.SetBasicAuth(service.username, service.password)
}

func (service *Service) ctx() (context.Context, func()) {
	return context.WithTimeout(context.Background(), time.Second*10)
}
