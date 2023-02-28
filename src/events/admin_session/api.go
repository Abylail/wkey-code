package admin_session

import (
	"bytes"
	"fmt"
	"net/http"
	"wkey-core/src/data/entities"
)

func (event *Event) Get(token string) (*entities.AdminSessionGet, error) {

	return nil, nil
}

// Active Активна ли сессия по токену
func (event *Event) Active(token string) (bool, error) {
	apiUrl := "http://127.0.0.1:8084/admin/api/v1/auth/login/token"
	authBody := []byte(fmt.Sprintf(`{"token": "%s"}`, token))

	request, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(authBody))
	if err != nil {
		return false, nil
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := http.DefaultClient
	response, err := client.Do(request)

	if err == nil && response.StatusCode == 200 {
		return true, nil
	}

	return false, nil
}
