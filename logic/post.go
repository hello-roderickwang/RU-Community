package logic

import (
	"web_app/dao/mysql"
	"web_app/models"
	"web_app/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//p.ID = int64(snowflake.GenID())
	p.ID = snowflake.GenID()
	return mysql.CreatePost(p)
}
