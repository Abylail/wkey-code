package events

import "wkey-core/src/events/migrate_event"

type ApiEvents struct {
	Migrate *migrate_event.Event
}

func Get() (*ApiEvents, error) {
	//config := definition.Config

	//rabbitConnection, err := rabbitmq_helper.Connection(config.Rabbit.ConnectionURL)
	//if err != nil {
	//	return nil, err
	//}

	migrate := migrate_event.Create()

	return &ApiEvents{
		Migrate: migrate,
	}, nil
}
