package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// ParamsSignUp 注册数据
type ParamsSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamsLogin 登录数据
type ParamsLogin struct {
	Username string `json:"username" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"` // 密码
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`       // 帖子id
	Direction int8   `json:"direction" binding:"oneof=-1 0 1"` // 赞成(1)反对(-1)取消投票(0)
}

// ParamPostList 帖子列表请求参数
type ParamPostList struct {
	Page        int64  `json:"page" form:"page"`                 // 页面数量
	Size        int64  `json:"size" form:"size"`                 // 页面大小
	CommunityID int64  `json:"community_id" form:"community_id"` // 可以为空
	Order       string `json:"order" form:"order"`               // 排序依据(time or score)
}

type ParamNewPost struct {
	Title       string `json:"title" binding:"required"`        // 标题
	Content     string `json:"content" binding:"required"`      // 内容
	CommunityID int64  `json:"community_id" binding:"required"` //所属社区id
}
