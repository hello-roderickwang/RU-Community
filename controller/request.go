package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("USER NOT LOGIN")

func getCurrentUserID(c *gin.Context) (userID int64, err error) {
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

func getPageInfo(c *gin.Context) (int64, int64) {
	pageStr := c.Query("page") // offset number, starting 1
	sizeStr := c.Query("size") // number of posts
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
