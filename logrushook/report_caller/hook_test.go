package report_caller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestHook(t *testing.T) {
	var (
		key    = "key"
		value  = "report_caller/hook_test.go:38"
		buffer bytes.Buffer
		fields logrus.Fields
	)

	hook := NewHook([]logrus.Level{logrus.ErrorLevel}, key)
	hook.SetLocationHandler(func(fileAbsolutePath string, line int) string {
		path := fileAbsolutePath
		s := "logrushook"
		index := strings.Index(fileAbsolutePath, s)
		if index != 0 {
			path = fileAbsolutePath[index+len(s)+1:]
		}
		return fmt.Sprintf("%s:%d", path, line)
	})

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)
	log.AddHook(hook)

	log.Error("message")
	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, value, fields[key])
}
