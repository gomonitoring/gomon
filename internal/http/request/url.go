package request

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Url struct {
	Name      string `json:"name"`
	Url       string `json:"url"`
	Threshold string `json:"threshold"`
}

type Stats struct {
	Name string `json:"name"`
}

func (u Url) Validate() error {
	if err := validation.ValidateStruct(&u,
		validation.Field(&u.Name, validation.Required, is.UTFLetterNumeric),
		validation.Field(&u.Url, validation.Required, is.URL),
		validation.Field(&u.Threshold, validation.Required, is.UTFNumeric),
	); err != nil {
		return fmt.Errorf("url validation failed %w", err)
	}

	return nil
}

func (s Stats) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.Name, validation.Required, is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("stats validation failed %w", err)
	}

	return nil
}
