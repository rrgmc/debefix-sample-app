package validatordeps

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func init() {
	enLocale := en.New()
	uni = ut.New(enLocale, enLocale)

	trans, _ := uni.GetTranslator("en")

	Validate = validator.New(validator.WithRequiredStructEnabled())
	_ = en_translations.RegisterDefaultTranslations(Validate, trans)
}
