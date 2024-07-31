package models

import "time"

type Post struct {
	Title   string `json:"title" db:"title" binding:"required"`
	Content string `json:"content" db:"content" binding:"required"`
	// 帖子id，通过雪花算法生成
	ID int64 `json:"id" db:"post_id,string"`
	// 用户ID，登录时保存到JWT中，通过解析JWT将ID放入ctx中
	AuthorID    int64     `json:"author_id,string" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

type ApiPostDetail struct {
	AuthorName       string                   `json:"author_name"`
	VoteNum          int64                    `json:"vote_num"`
	*Post                                     // 嵌入帖子结构体
	*CommunityDetail `json:"CommunityDetail"` // 嵌入社区信息结构体
}
