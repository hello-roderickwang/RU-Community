package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("User Exist")
	ErrorUserNotExist    = errors.New("User Not Exist")
	ErrorInvalidPassword = errors.New("Wrong Password")
	ErrorInvaldID        = errors.New("Invalid ID")
)
