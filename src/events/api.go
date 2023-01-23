package events

type ApiEvents struct {
	//
}

func Get() (*ApiEvents, error) {
	return &ApiEvents{}, nil
}
