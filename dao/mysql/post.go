package mysql

import (
	"bluebell/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return
}

// GetPostDetailByID 根据ID查询帖子详情
func GetPostDetailByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
			   from post
			   where post_id = ?`
	if err = db.Get(post, sqlStr, id); err != nil {
		zap.L().Error("db.Get(post, sqlStr, id)", zap.Error(err))
		return nil, err
	}
	return
}

func GetUserById(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id, username from user where user_id = ?`
	if err = db.Get(user, sqlStr, uid); err != nil {
		zap.L().Error("db.Get(user, sqlStr, uid)", zap.Error(err))
		return nil, err
	}
	return
}

func GetCommunityById(cid int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	fmt.Printf("community id = %v\n", cid)
	sqlStr := `select community_name, introduction from community where community_id = ?`
	if err = db.Get(community, sqlStr, cid); err != nil {
		zap.L().Error("db.Get(community, sqlStr, cid)", zap.Error(err))
		return nil, err
	}
	fmt.Printf("community: %v\n", community)
	return
}

// GetPostList 获取帖子列表，分页展示
func GetPostList(page, size int64) (postList []*models.Post, err error) {
	sqlStr := `select 
    		   post_id, title, content, author_id, community_id, create_time
			   from post
			   Order By create_time
			   DESC 
			   limit ?,?
			   `

	postList = make([]*models.Post, 0, 2)                   // 长度、容量
	err = db.Select(&postList, sqlStr, (page-1)*size, size) // 起始数据位置、页面大小
	if err != nil {
		zap.L().Error(" db.Select(&postList, sqlStr) error", zap.Error(err))
		return nil, err
	}
	return

}

// GetPostListByIDs 根据id列表去mysql查询post详细信息
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id, ?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		zap.L().Error("sqlx.In(sqlStr, ids, strings.Join(ids, error", zap.Error(err))
		return nil, err
	}
	query = db.Rebind(query)                   // 防止sql注入
	err = db.Select(&postList, query, args...) // args...展开参数
	return
}
