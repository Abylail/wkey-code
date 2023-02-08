package middlewares

import (
	"wkey-core/src/events/admin_session"
	"wkey-core/src/events/client_session"
)

func Init() {
	adminSession = admin_session.Create()
	clientSession = client_session.Create()
}
