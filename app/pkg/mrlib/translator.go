package mrlib

import (
    "auth-le-back/pkg/mrapp"
    "fmt"

    "github.com/ilyakaznacheev/cleanenv"
)

type (
    langMap map[mrapp.LangCode]*Locale

    Locale struct {
        code mrapp.LangCode
        Messages map[mrapp.MessageCode]string `yaml:"messages"`
        Errors map[mrapp.ErrorCode]mrapp.ErrorMessage `yaml:"errors"`
    }

    Translator struct {
        logger mrapp.Logger
        langs langMap
        defaultLocale *Locale
    }

    TranslatorOptions struct {
        DirPath string
        FileType string
        LangCodes mrapp.LangCodes
    }
)

// Make sure the Translator conforms with the mrapp.Translator interface
var _ mrapp.Translator = (*Translator)(nil)

func NewTranslator(logger mrapp.Logger, opt TranslatorOptions) *Translator {
    var defaultLocale *Locale
    langs := langMap{}

    if len(opt.LangCodes) == 0 {
        logger.Warn("opt.LangCodes is empty")
    }

    for i, code := range opt.LangCodes {
        locale, err := newLocale(code, fmt.Sprintf("%s/%s.%s", opt.DirPath, code, opt.FileType))

        if err != nil {
            logger.Error(err)
            continue
        }

        langs[code] = locale

        if i == 0 {
            defaultLocale = locale
        }
    }

    return &Translator{
        logger: logger,
        langs: langs,
        defaultLocale: defaultLocale,
    }
}

func newLocale(code mrapp.LangCode, filePath string) (*Locale, error) {
    locale := &Locale{
        code: code,
    }

    if err := cleanenv.ReadConfig(filePath, locale); err != nil {
        return nil, fmt.Errorf("while reading locale '%s', error '%s' occurred", filePath, err)
    }

    return locale, nil
}

func (t Translator) GetLocale(langs ...mrapp.LangCode) mrapp.Locale {
    for _, lang := range langs {
        if locale, ok := t.langs[lang]; ok {
            return locale
        }
    }

    if t.defaultLocale != nil {
        return t.defaultLocale
    }

    return &Locale{} // stub
}

func (t Translator) GetLocaleByAcceptLanguage(s string) mrapp.Locale {
    return t.GetLocale(ParseAcceptLanguage(s)...)
}

func (l Locale) GetCode() mrapp.LangCode {
    return l.code
}

func (l Locale) GetError(errorCode mrapp.ErrorCode) mrapp.ErrorMessage {
    value, ok := l.Errors[errorCode]

    if ok {
        return value
    }

    return mrapp.ErrorMessage{Reason: string(errorCode)} // stub
}
