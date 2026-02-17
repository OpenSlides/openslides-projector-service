package i18n

import (
	"github.com/leonelquinteros/gotext"
	"golang.org/x/text/language"
)

type ProjectorLocale struct {
	locale             *gotext.Locale
	customTranslations map[string]string
}

func NewLocale(lang language.Tag) *ProjectorLocale {
	langName, _ := lang.Base()
	locale := gotext.NewLocale("locale", langName.String())
	locale.AddDomain("default")

	return &ProjectorLocale{
		locale: locale,
	}
}

func (p *ProjectorLocale) Get(str string, vars ...any) string {
	str = p.locale.Get(str, vars...)

	if custom, ok := p.customTranslations[str]; ok {
		return custom
	}

	return str
}

func (p *ProjectorLocale) SetCustomTranslations(translation map[string]string) {
	p.customTranslations = translation
}
