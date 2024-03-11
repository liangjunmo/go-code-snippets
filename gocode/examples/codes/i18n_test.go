package codes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTranslate(t *testing.T) {
	message := Translate("code", "")
	require.Equal(t, zhCn[Unknown], message)
}
