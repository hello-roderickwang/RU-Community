package mysql

import "web_app/models"

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id, title, content, author_id, community_id)
	values (?, ?, ?, ?, ?)
	`
	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

func GetPostByID(pid int64) (post *models.Post, err error) {
	sqlStr := `select
		post_id, title, content, author_id, community_id, create_time
		from post
		where post_id = ?`
	post = new(models.Post)
	err = db.Get(post, sqlStr, pid)
	return
}
