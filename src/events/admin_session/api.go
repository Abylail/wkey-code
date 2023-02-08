package admin_session

import "wkey-core/src/data/entities"

func (event *Event) Get(token string) (*entities.AdminSessionGet, error) {
	return nil, nil
}

func (event *Event) Active(token string) (bool, error) {
	return false, nil
}
