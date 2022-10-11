package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Email string

func NewEmail(email string) (Email, error) {
	e := Email(email)
	if err := e.validate(); err != nil {
		return "", err
	}
	return e, nil
}

func (e *Email) validate() error {
	return validation.Validate(e, is.Email)
}

func (e Email) String() string {
	return string(e)
}

func (e Email) IsEmpty() bool {
	return e == ""
}
