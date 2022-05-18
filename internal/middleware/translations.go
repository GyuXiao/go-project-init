// Package middleware
/**
  @author: zk.xiao
  @date: 2022/5/18
  @note:
**/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	validator "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// Translations
/**
 * @Description: 实现对 validator 的语言包翻译的相关功能（正确地调包就能解决问题）
 * @return gin.HandlerFunc
 */
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())

		// 识别当前请求的语言类别
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)

		if ok {
			// 调用 RegisterDefaultTranslations 方法将验证器和对应语言类型的 Translator 注册进来，实现验证器的多语言支持，
			switch locale {
			case "zh":
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = enTranslations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zhTranslations.RegisterDefaultTranslations(v, trans)
				break
			}
			// 将 Translator 存储到全局上下文中，便于后续翻译时的使用
			c.Set("trans", trans)
		}
		c.Next()
	}
}
