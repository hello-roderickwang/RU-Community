package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// check existing user
	mysql.QueryUserByUserName()
	// generate user id
	snowflake.GenID()
	// save data to database
	mysql.InsertUser()

}
