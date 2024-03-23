package report_caller

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestHook(t *testing.T) {
	key := "key"
	value := "value"

	hook := NewHook([]logrus.Level{logrus.ErrorLevel}, key)
	hook.SetLocationHandler(func(fileAbsolutePath string, line int) string {
		return value
	})

	var buffer bytes.Buffer
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)
	log.AddHook(hook)

	log.Error("message")
	var fields logrus.Fields
	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, value, fields[key])
}
