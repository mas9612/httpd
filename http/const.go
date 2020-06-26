package http

const (
	StatusInvalidRequest       = 400
	StatusMethodNotImplemented = 501
)

var (
	supportedMethods = []string{
		"HEAD",
		"GET",
	}
)
