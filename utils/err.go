package utils

import "errors"

// concat multiple strings/errors into a single error
func Anyhow(errs ...interface{}) error {
	erros := []error{}
	for _, err := range errs {
		switch err := err.(type) {
		case string:
			erros = append(erros, errors.New(err))
		case error:
			erros = append(erros, err)
		default:
			panic("Invalid type for err")
		}
	}
	return errors.Join(erros...)
}
