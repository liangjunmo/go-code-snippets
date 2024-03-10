package codes

import (
	"errors"

	"github.com/liangjunmo/goutil/gocode"
)

func Parse(err error, language Language) (code gocode.Code, message string) {
	code = gocode.Parse(err)

	if errors.Is(code, gocode.DefaultCode) {
		code = Unknown
	} else if errors.Is(code, gocode.SuccessCode) {
		code = OK
	}

	message = Translate(code, language)

	return code, message
}
