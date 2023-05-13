package mrapp

type (
    LangCode string // ISO 639 and regions
    LangCodes []LangCode
    MessageCode string
    ErrorCode string

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

func (l LangCodes) Convert(langs ...string) LangCodes {
    var langCodes LangCodes

    for _, lang := range langs {
        langCodes = append(langCodes, LangCode(lang))
    }

    return langCodes
}
