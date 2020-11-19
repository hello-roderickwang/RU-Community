package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("USER NOT LOGIN")

func getCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
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
