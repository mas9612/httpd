package http

const (
	// StatusOK represents the status code 200.
	StatusOK = 200
	// StatusBadRequest represents the status code 400.
	StatusBadRequest = 400
	// StatusNotFound represents the status code 404.
	StatusNotFound = 404
	// StatusInternalServerError represents the status code 500.
	StatusInternalServerError = 500
	// StatusMethodNotImplemented represents the status code 501.
	StatusMethodNotImplemented = 501
)

var reasonPhrases = map[int]string{
	StatusOK:                   "OK",
	StatusBadRequest:           "Bad Request",
	StatusNotFound:             "Not Found",
	StatusInternalServerError:  "InternalServerError",
	StatusMethodNotImplemented: "Not Implemented",
}
