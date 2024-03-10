package gocode

const (
	DefaultCode Code = "default"
	SuccessCode Code = "success"
)

type Code string

func (code Code) Error() string {
	return string(code)
}

func (code Code) String() string {
	return string(code)
}
