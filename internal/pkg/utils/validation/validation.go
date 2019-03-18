package validation

import (
	ozzo "github.com/go-ozzo/ozzo-validation"
	"github.com/zeyneloz/sample-go-rest-api/internal/pkg/errors"
)

// Schema describes the fields of a post request.
type Schema interface {
	Validate() error
}

// Validate schema and return http error
func Validate(schema Schema, rules ...*ozzo.FieldRules) error {
	err := ozzo.ValidateStruct(schema, rules...)
	switch e := err.(type) {
	case nil:
		return nil
	case ozzo.Errors:
		// Form a proper HTTPError
		messages := make(map[string]string)
		for k, v := range e {
			messages[k] = v.Error()
		}
		return errors.NewHTTPValidationError(messages)
	default:
		return errors.NewInternalError(e, "Error while validating")
	}
}
