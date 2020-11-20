package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"web_app/logic"
	"web_app/models"
)

func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := errs.Translate(trans)
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	logic.PostVote()
	ResponseSuccess(c, nil)
}
