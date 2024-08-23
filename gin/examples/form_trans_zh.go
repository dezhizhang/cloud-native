package examples

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	zh_translation "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

type UserForm struct {
	Name     string `form:"name" binding:"required,min=2,max=20"`
	Password string `form:"password" binding:"required,min=8,max=20"`
}

var Trans ut.Translator

func translation(locale string) (err error) {
	var ok bool
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文鄱译器
	uni := ut.New(zhT, enT)
	v := validator.New()
	Trans, ok = uni.GetTranslator(locale)

	if !ok {
		return errors.New("uni.GetTranslator failed")
	}

	switch locale {
	case "en":
		_ = en_translation.RegisterDefaultTranslations(v, Trans)
		break
	case "zh":
		_ = zh_translation.RegisterDefaultTranslations(v, Trans)
		break
	default:
		_ = en_translation.RegisterDefaultTranslations(v, Trans)
	}
	return err

}

func handleLoginZh(c *gin.Context) {
	var user UserForm
	if err := c.ShouldBind(&user); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": errs.Error()})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errs.Translate(Trans)})
		return
	}
	c.JSON(http.StatusOK, &user)
}

func main() {
	err := translation("zh")
	if err != nil {
		fmt.Println(err)
	}
	r := gin.Default()
	r.POST("/login", handleLoginZh)
	_ = r.Run(":8080")
}
