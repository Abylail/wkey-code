package router

type Config struct {
	Method   string
	Endpoint string
	Port     string

	Body    any
	Headers map[string][]string
	Cookies map[string]string
}
