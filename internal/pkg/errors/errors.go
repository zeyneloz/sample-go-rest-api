package errors

import (
	"encoding/json"
	"net/http"

	pkgerrors "github.com/pkg/errors"
)

// BaseError is a trivial implementation of error.
type BaseError struct {
	detail string
}

func (e *BaseError) Error() string {
	return e.detail
}

// New returns an error that formats as the given text.
func New(detail string) error {
	return pkgerrors.WithStack(&BaseError{detail})
}

// UserVisibleError is an error whose details to be shared with user.
type UserVisibleError interface {
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}

// InternalError is an error caused by another error.
type InternalError struct {
	detail string
	cause  error
}

func (e *InternalError) Error() string {
	return e.detail + "; " + e.cause.Error()
}

// NewInternalError returns an error with given cause error and detail.
func NewInternalError(cause error, detail string) error {
	return pkgerrors.WithStack(&InternalError{detail, cause})
}

// ResponseBody returns empty body.
func (e *InternalError) ResponseBody() ([]byte, error) {
	return []byte{}, nil
}

// ResponseHeaders returns http status code and headers.
func (e *InternalError) ResponseHeaders() (int, map[string]string) {
	return http.StatusInternalServerError, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

// HTTPError is handled by a middleware and returned in response.
type HTTPError struct {
	Detail   string            `json:"detail"`
	Messages map[string]string `json:"errors"`
	Status   int               `json:"-"`
}

func (e *HTTPError) Error() string {
	return e.Detail
}

// ResponseBody returns response body.
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)
	if err != nil {
		return nil, NewInternalError(err, "Error while parsing json")
	}
	return body, nil
}

// ResponseHeaders returns http status code and headers.
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

// NewHTTPError returns a new user visible http error.
func NewHTTPError(status int, detail string, messages map[string]string) error {
	return pkgerrors.WithStack(&HTTPError{
		Detail:   detail,
		Status:   status,
		Messages: messages,
	})
}

// NewHTTPValidationError returns HTTPError with BadRequest status.
func NewHTTPValidationError(messages map[string]string) error {
	return NewHTTPError(http.StatusBadRequest, "Validation error", messages)
}

// Cause returns root cause of error.
func Cause(err error) error {
	return pkgerrors.Cause(err)
}
