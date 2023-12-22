package mysql

import (
	"goweb/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`

	_, err = db.Exec(sqlStr, p.ID, p.Title, p.Content, p.AuthorID, p.CommunityID)
	return
}

// GetPostById 根据id查询单个贴子数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	post_id, title, content, author_id, community_id, create_time
	from post
	where post_id = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}

func GetUserByID(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`

	err = db.Get(user, sqlStr, uid)

	return user, err
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
	post_id, title, content, author_id, community_id, create_time
	from post
	ORDER BY create_time
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
