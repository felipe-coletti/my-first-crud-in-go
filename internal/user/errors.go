package user

import (
	"encoding/json"
	"errors"

	rest_err "github.com/felipe-coletti/my-first-crud-in-go/pkg/errors"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	translate ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)

		translate, _ = un.GetTranslator("en")

		en_translation.RegisterDefaultTranslations(val, translate)
	}
}

func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewRestErr("bad_request", "Invalid field type", nil)
	} else if errors.As(validation_err, &jsonValidationError) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(translate),
				Field: e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}

		return rest_err.NewRestErr("bad_request", "Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewRestErr("bad_request", "Error trying to convert field", nil)
	}
}