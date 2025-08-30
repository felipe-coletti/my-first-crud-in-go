package errors

import "net/http"

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

func (r *RestErr) Error() string {
	return r.Message
}

var errorTypes = map[string]struct {
	code int
	err  string
}{
	"bad_request": {
		code: http.StatusBadRequest,
		err:  "bad_request",
	},
	"internal_server_error": {
		code: http.StatusInternalServerError,
		err:  "internal_server_error",
	},
	"not_found": {
		code: http.StatusNotFound,
		err:  "not_found",
	},
	"forbidden": {
		code: http.StatusForbidden,
		err:  "forbidden",
	},
}

func NewRestErr(errorType, message string, causes []Causes) *RestErr {
	if et, ok := errorTypes[errorType]; ok {
		return &RestErr{
			Message: message,
			Err:     et.err,
			Code:    et.code,
			Causes:  causes,
		}
	}

	return &RestErr{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
		Causes:  causes,
	}
}
