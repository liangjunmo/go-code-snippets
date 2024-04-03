package codes

import (
	"errors"

	"github.com/liangjunmo/go-code-snippets/code"
)

func Parse(err error) code.Code {
	c := code.Parse(err)
	if errors.Is(c, code.DefaultCode) {
		c = Unknown
	} else if errors.Is(c, code.SuccessCode) {
		c = OK
	}
	return c
}
