package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"web_app/middlewares"
)

var ErrorUserNotLogin = errors.New("USER NOT LOGIN")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(middlewares.CtxUserIDKey)
	//if !ok {
	//	err = ErrorUserNotLogin
	//	// what is the userID for this return
	//	return
	//}
	//uid, ok = uid.(int64)
	//if !ok {
	//	err = ErrorUserNotLogin
	//	return
	//}

	if ok {
		userID, ok = uid.(int64)
	}
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}
