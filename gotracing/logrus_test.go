package gotracing

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogrusHook(t *testing.T) {
	resetTracingKeys()

	key := "key"
	value := "value"

	SetTracingIDKey(key)
	SetTracingIDGenerator(func() string {
		return value
	})

	var (
		buffer bytes.Buffer
		fields logrus.Fields
	)
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)
	log.AddHook(NewLogrusHook())

	ctx := Trace(context.Background())
	log.WithContext(ctx).Error("message")
	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, value, fields[key])
}
