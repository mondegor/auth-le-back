package mrlib

import (
    "auth-le-back/pkg/mrapp"
    "regexp"
    "strings"
)

const (
    acceptLanguageByDefault = "en"
)

var regexpAcceptLanguage = regexp.MustCompile("^[a-z]{2}(-[a-zA-Z0-9-]+)?$")

// ParseAcceptLanguage
// Sample Accept-Language: ru;q=0.9, fr-CH, fr;q=0.8, en;q=0.7, *;q=0.5
func ParseAcceptLanguage(s string) []mrapp.LangCode {
    length := len(s)

    if length > 0 && length <= 128 {
        langs := mrapp.LangCodes{}
        keys := make(map[string]bool)

        for _, lang := range strings.Split(strings.ToLower(s), ",") {
            if index := strings.Index(lang, ";"); index >= 0 {
                lang = lang[:index]
            }

            lang = strings.TrimSpace(lang)

            if !regexpAcceptLanguage.MatchString(lang) {
                continue
            }

            if _, isExists := keys[lang]; isExists {
                continue
            }

            langs = append(langs, mrapp.LangCode(lang))
            keys[lang] = true
        }

        if len(langs) > 0 {
            return langs
        }
    }

    return mrapp.LangCodes{acceptLanguageByDefault}
}
