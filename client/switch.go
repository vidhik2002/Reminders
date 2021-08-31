package client

type BackendHTTPClient interface {
}
type Switch struct {
	client BackendHTTPClient
	backendAPIURL string
	commands map[string] func() func(string) error
}

func NewSwitch(uri string) Switch {
	httpClient := NewHTTPClient(uri)
	s := Switch{
		client: httpClient,
		backendAPIURL: uri,
	}
	return s
}