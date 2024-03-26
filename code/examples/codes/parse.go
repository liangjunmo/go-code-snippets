package codes

import (
	"errors"

	"github.com/liangjunmo/goutil/codes"
)

func Parse(err error) codes.Code {
	code := codes.Parse(err)
	if errors.Is(code, codes.DefaultCode) {
		code = Unknown
	} else if errors.Is(code, codes.SuccessCode) {
		code = OK
	}
	return code
}
