package request

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	minUsernameLen = 3
	maxUsernameLen = 10
	minPassLen     = 8
	maxPassLen     = 20
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) Validate() error {
	if err := validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(minUsernameLen, maxUsernameLen), is.UTFLetterNumeric),
		validation.Field(&u.Password, validation.Required, validation.Length(minPassLen, maxPassLen), is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("user validation failed %w", err)
	}

	return nil
}
