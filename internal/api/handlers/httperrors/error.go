package httperrors

type HTTPError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(message string, statusCode int) HTTPError {
	return HTTPError{
		Message:    message,
		StatusCode: statusCode,
	}
}
