package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"web_app/models"
)

const secrete = "roderickwang"

func CheckUserExist(username string) error {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if count > 0 {
		return errors.New("USER ALREADY EXIST")
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
