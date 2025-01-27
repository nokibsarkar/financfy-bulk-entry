package Exception

import "net/http"

type Failure struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func UnexpectedServerFailure(message string) *Failure {
	return &Failure{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

func NotFoundError(message string) *Failure {
	return &Failure{
		StatusCode: http.StatusNotFound,
		Message:    message,
	}
}

func BadRequestError(message string) *Failure {
	return &Failure{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}
