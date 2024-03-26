package code

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse_Nil(t *testing.T) {
	require.Equal(t, SuccessCode, Parse(nil))
}

func TestParse_Error(t *testing.T) {
	err := fmt.Errorf("err")
	require.Equal(t, DefaultCode, Parse(err))
}

func TestParse_WrappedError(t *testing.T) {
	err := fmt.Errorf("err")
	err = fmt.Errorf("%s: %w", "err", err)
	require.Equal(t, DefaultCode, Parse(err))
}

func TestParse_Code(t *testing.T) {
	code := Code("code")
	require.Equal(t, code, Parse(code))
}

func TestParse_WrappedCode(t *testing.T) {
	code := Code("code")
	err := fmt.Errorf("%s: %w", "err", code)
	require.Equal(t, code, Parse(err))
}

func TestParse_MultipleWrappedCode(t *testing.T) {
	code := Code("code")
	code10 := Code("code10")
	err := fmt.Errorf("%s: %w", "err", code)
	for i := 1; i < 10; i++ {
		c := fmt.Sprintf("code%d", i)
		err = fmt.Errorf("%s: %w", "err", Code(c))
	}
	err = fmt.Errorf("%s: %w", "err", code10)
	require.Equal(t, code10, Parse(err))
}
