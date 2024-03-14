package codes

import (
	"errors"

	"github.com/liangjunmo/goutil/gocode"
)

func Parse(err error) gocode.Code {
	code := gocode.Parse(err)
	if errors.Is(code, gocode.DefaultCode) {
		code = Unknown
	} else if errors.Is(code, gocode.SuccessCode) {
		code = OK
	}
	return code
}
