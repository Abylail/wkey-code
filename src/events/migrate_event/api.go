package migrate_event

import (
	"net/http"
	"wkey-core/src/services/requests"
)

func (event *Event) Categories() error {
	_, status, err := requests.New(http.MethodPost, categoriesURL, nil).SendStatus()
	if err != nil {
		return err
	}

	if status == http.StatusUnauthorized {
		// TODO: что-то сделать с авторизацией (ее пока нет)
	}

	return nil
}

func (event *Event) Products() error {
	_, status, err := requests.New(http.MethodPost, productsURL, nil).SendStatus()
	if err != nil {
		return err
	}

	if status == http.StatusUnauthorized {
		// TODO: что-то сделать с авторизацией (ее пока нет)
	}

	return nil
}
