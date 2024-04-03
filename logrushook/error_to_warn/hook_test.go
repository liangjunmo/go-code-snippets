package error_to_warn

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"

	"github.com/liangjunmo/go-code-snippets/code"
)

func TestHook_TransformToWarn(t *testing.T) {
	var (
		buffer   bytes.Buffer
		fields   logrus.Fields
		notFound code.Code = "NotFound"
		key                = "level"
		value              = "warning"
	)

	hook := NewHook(nil, true)

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)
	log.AddHook(hook)

	log.WithError(notFound).Error("message")
	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, value, fields[key])
}

func TestHook_NotTransformToWarn(t *testing.T) {
	var (
		buffer              bytes.Buffer
		fields              logrus.Fields
		internalServerError code.Code = "InternalServerError"
		key                           = "level"
		value                         = "error"
	)

	hook := NewHook([]code.Code{internalServerError}, true)

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)
	log.AddHook(hook)

	log.WithError(internalServerError).Error("message")
	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, value, fields[key])
}
