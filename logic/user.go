package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/jwt"
	"web_app/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// check existing user
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		return err
	}

	// generate user id
	userID := snowflake.GenID()

	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}

	// save data to database
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// user is a pointer
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	return jwt.GenToken(user.UserID, user.Username)
}
