package gocode

import (
	"errors"
)

func Parse(err error) Code {
	if err == nil {
		return SuccessCode
	}

	code := DefaultCode

	for {
		if err == nil {
			break
		}

		if errors.As(err, &code) {
			break
		}

		err = errors.Unwrap(err)
	}

	return code
}
