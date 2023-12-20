package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"goweb/logic"
	"goweb/models"
	"net/http"
)

func SignUpHandler(c *gin.Context) {

	// check format

	p := new(models.ParamSignUp)

	if err := c.ShouldBindJSON(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"password":   p.Password,
			"repassword": p.RePassword,
		})
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// json format err
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return
	}

	if err := logic.SignUp(p); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "SignUp failed",
		})
		return
	}

	// call back
	c.JSON(http.StatusOK, gin.H{
		"msg": "SignUp success",
	})

	return
}

//func LoginHandler(c *gin.Context) {
//	p := new(models.ParamLogin)
//	if err := c.ShouldBindJSON(p); err != nil {
//		zap.L().Error("Login with invalid param", zap.Error(err))
//
//		errs, ok := err.(validator.ValidationErrors)
//
//		if !ok {
//			c.JSON(http.StatusOK, gin.H{
//
//			})
//		}
//	}
//}
