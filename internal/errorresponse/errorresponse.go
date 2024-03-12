package errorresponse

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type ErrorResponse struct {
	HttpStatusCode int    `json:"-"`
	Message        string `json:"message"`
}

var (
	RequestTimeout = ErrorResponse{
		HttpStatusCode: http.StatusRequestTimeout,
		Message:        http.StatusText(http.StatusRequestTimeout),
	}
	ValidationFailed = ErrorResponse{
		HttpStatusCode: http.StatusBadRequest,
		Message:        http.StatusText(http.StatusBadRequest) + " - invalid json body",
	}
)

func (e ErrorResponse) Error() string {
	return e.Message
}

func WrapErrorWithStatus(err error, status int) ErrorResponse {
	if e, ok := err.(ErrorResponse); ok {
		return e
	}

	return ErrorResponse{
		HttpStatusCode: status,
		Message:        errors.Wrap(err, http.StatusText(status)).Error(),
	}
}

func WrapError(err error) ErrorResponse {
	if e, ok := err.(ErrorResponse); ok {
		return e
	}

	switch err {
	case context.DeadlineExceeded:
		return RequestTimeout
	case context.Canceled:
		// Verify if context is cancelled due to db query timing out
		// or if client cancelled request
		if err.Error() == "pq: canceling statement due to user request" {
			return InternalServerError(err)
		} else {
			return ErrorResponse{
				HttpStatusCode: http.StatusBadRequest,
				Message:        "request cancelled",
			}
		}
	default:
		return InternalServerError(err)
	}
}

func NewError(msg string, status int) ErrorResponse {
	return ErrorResponse{
		HttpStatusCode: status,
		Message:        msg,
	}
}

func BadRequest(err error) ErrorResponse {
	return WrapErrorWithStatus(err, http.StatusBadRequest)
}

func InternalServerError(err error) ErrorResponse {
	return WrapErrorWithStatus(err, http.StatusInternalServerError)
}
