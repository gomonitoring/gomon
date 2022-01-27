package request

import (
	"fmt"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Alert struct {
	UrlName string `json:"urlname"`
}

func (a Alert) Validate() error {
	if err := validation.ValidateStruct(&a,
		validation.Field(&a.UrlName, validation.Required, is.UTFLetterNumeric),
	); err != nil {
		return fmt.Errorf("alert validation failed %w", err)
	}

	return nil
}
