package middleware

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni          *ut.UniversalTranslator
	validate     *validator.Validate
	zhTrans      ut.Translator
	enTrans      ut.Translator
	registerOnce sync.Once
)

func init() {
	uni = ut.New(en.New(), zh.New(), zh_Hant_TW.New())
	validate, _ = binding.Validator.Engine().(*validator.Validate)

	zhTrans, _ = uni.GetTranslator("zh")
	enTrans, _ = uni.GetTranslator("en")

	// Register translations once
	registerOnce.Do(func() {
		_ = zh_translations.RegisterDefaultTranslations(validate, zhTrans)
		_ = en_translations.RegisterDefaultTranslations(validate, enTrans)
	})
}

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.GetHeader("locale")
		var trans ut.Translator

		switch locale {
		case "en":
			trans = enTrans
		case "zh":
			trans = zhTrans
		default:
			trans = zhTrans
		}

		c.Set("trans", trans)
		c.Next()
	}
}
