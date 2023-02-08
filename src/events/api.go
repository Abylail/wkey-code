package events

import (
	"wkey-core/src/events/admin_session"
	"wkey-core/src/events/client_session"
	"wkey-core/src/events/migrate_event"
)

type ApiEvents struct {
	Migrate *migrate_event.Event

	AdminSession  *admin_session.Event
	ClientSession *client_session.Event
}

func Get() (*ApiEvents, error) {
	//config := definition.Config

	//rabbitConnection, err := rabbitmq_helper.Connection(config.Rabbit.ConnectionURL)
	//if err != nil {
	//	return nil, err
	//}

	migrate := migrate_event.Create()

	adminSession := admin_session.Create()
	clientSession := client_session.Create()

	return &ApiEvents{
		Migrate:       migrate,
		AdminSession:  adminSession,
		ClientSession: clientSession,
	}, nil
}
