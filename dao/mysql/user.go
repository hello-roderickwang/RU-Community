package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web_app/models"
)

const secrete = "roderickwang"

func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return ErrorUserExist
	}
	return err
}

func InsertUser(user *models.User) (err error) {
	// encrypt password
	user.Password = encryptPassword(user.Password)
	// execute sql
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secrete))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	oPassword := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}
