package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web_app/logic"
	"web_app/models"
)

// function to deal with sign up request
func SignUpHandler(c *gin.Context) {
	// 1. get parameter and parameter test
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// validator
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans),
		})
		return
	}

	// manually validate param
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "request parameter error",
	//	})
	//	return
	//}

	fmt.Println(p)
	// 2. deal parameter
	logic.SignUp(p)
	// 3. return
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
	})
}
