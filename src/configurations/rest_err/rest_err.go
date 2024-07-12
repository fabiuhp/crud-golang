package rest_err

import "net/http"

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, nil)
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return NewRestErr(message, "bad_request", http.StatusBadRequest, causes)
}

func NewInternalServerError(message string) *RestErr {
	return NewRestErr(message, "internal_server_error", http.StatusInternalServerError, nil)
}

func NewNotFoundError(message string) *RestErr {
	return NewRestErr(message, "not_found", http.StatusNotFound, nil)
}

func NewForbiddenError(message string) *RestErr {
	return NewRestErr(message, "forbidden", http.StatusForbidden, nil)
}
