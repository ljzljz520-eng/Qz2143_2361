package policy

import (
	"community50/model"
	"errors"
	"strings"
)

func ValidateTitle(v string) error {
	v = strings.TrimSpace(v)
	if v == "" {
		return errors.New("title required")
	}
	if len([]rune(v)) > 80 {
		return errors.New("title too long")
	}
	return nil
}
func ValidateBody(v string) error {
	if strings.TrimSpace(v) == "" {
		return errors.New("body required")
	}
	return nil
}
func Validate(r model.Record) error {
	if e := ValidateTitle(r.Title); e != nil {
		return e
	}
	return ValidateBody(r.Body)
}
func IsPublic(r model.Record) bool { return r.Status == "immediate" }
