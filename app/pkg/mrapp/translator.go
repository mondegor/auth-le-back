package mrapp

import (
    "strings"
)

type (
    LangCode string // ISO 639 and regions
    LangCodes []LangCode
    MessageCode string

    ErrorMessage struct {
        Reason string `yaml:"reason"`
        Details []string `yaml:"details"`
    }
)

type Translator interface {
    GetLocale(langs ...LangCode) Locale
    GetLocaleByAcceptLanguage(s string) Locale
}

type Locale interface {
    GetCode() LangCode
    GetError(errorCode ErrorCode) ErrorMessage
}

func (em *ErrorMessage) GetDetails() string {
    switch len(em.Details) {
        case 0:
            return ""
        case 1:
            return em.Details[0]
    }

    return "- " + strings.Join(em.Details, "\n- ")
}

func CastToLangCodes(langs ...string) LangCodes {
    var langCodes LangCodes

    for _, lang := range langs {
        langCodes = append(langCodes, LangCode(lang))
    }

    return langCodes
}
