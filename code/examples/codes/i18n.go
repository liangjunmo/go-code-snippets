package codes

import (
	"github.com/liangjunmo/goutil/codes"
)

type Language string

const (
	LanguageZhCn Language = "zh_CN"
)

var i18n = map[Language]map[codes.Code]string{
	LanguageZhCn: zhCn,
}

func Translate(code codes.Code, language Language) string {
	if _, ok := i18n[language]; !ok {
		language = LanguageZhCn
	}
	if _, ok := i18n[language][code]; !ok {
		code = Unknown
	}
	return i18n[language][code]
}
