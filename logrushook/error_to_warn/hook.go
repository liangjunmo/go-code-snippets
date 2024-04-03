package error_to_warn

import (
	"context"
	"errors"

	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/go-code-snippets/code"
)

type Hook struct {
	codesNotToWarn []code.Code
	deleteErrorKey bool
}

func NewHook(codesNotToWarn []code.Code, deleteErrorKey bool) logrus.Hook {
	return &Hook{
		codesNotToWarn: codesNotToWarn,
		deleteErrorKey: deleteErrorKey,
	}
}

func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)
	if !ok || err == nil {
		return nil
	}

	if hook.deleteErrorKey {
		delete(entry.Data, logrus.ErrorKey)
	}

	if errors.Is(err, context.Canceled) {
		entry.Level = logrus.WarnLevel
		return nil
	}

	for _, c := range hook.codesNotToWarn {
		if errors.Is(err, c) {
			return nil
		}
	}
	entry.Level = logrus.WarnLevel
	return nil
}
