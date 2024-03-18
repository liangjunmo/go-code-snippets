package codes

import (
	"github.com/liangjunmo/goutil/gocode"
)

type Language string

const (
	LanguageZhCn Language = "zh_CN"
)

var i18n = map[Language]map[gocode.Code]string{
	LanguageZhCn: zhCn,
}

func Translate(code gocode.Code, language Language) string {
	if _, ok := i18n[language]; !ok {
		language = LanguageZhCn
	}
	if _, ok := i18n[language][code]; !ok {
		code = Unknown
	}
	return i18n[language][code]
}
