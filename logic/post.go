package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"go.uber.org/zap"
)

// CreatePost 创建帖子
func CreatePost(post *models.Post) (err error) {
	// 1. 生成post ID
	post.ID = snowflake.GenID()

	// 2. 保存到数据库
	// mysql创建帖子
	if err := mysql.CreatePost(post); err != nil {
		zap.L().Error("mysql.CreatePost(&post) failed", zap.Error(err))
		return err
	}
	// redis创建帖子
	if err := redis.CreatePost(post.ID, post.CommunityID); err != nil {
		zap.L().Error("redis.CreatePost(&post) failed", zap.Error(err))
		return err
	}

	// 3. 返回
	return
}

// GetPostDetailByID 根据ID查询帖子详情
func GetPostDetailByID(id int64) (apiPostDetail *models.ApiPostDetail, err error) {
	// 调用dao层查询数据并拼接
	// 查询post
	post, err := mysql.GetPostDetailByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostDetailByID(id) failed", zap.Error(err))
		return nil, err
	}

	// 查询authorName(根据post的author_id)
	user, err := mysql.GetUserById(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("AuthorId:", post.AuthorID),
			zap.Error(err))
		return nil, err
	}

	// 查询Community(根据post的community_id)
	communityDetail, err := mysql.GetCommunityById(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityById(post.CommunityID) failed", zap.Error(err))
		return nil, err
	}

	// 拼接ApiPostDetail
	apiPostDetail = &models.ApiPostDetail{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: communityDetail,
	}

	return
}

// GetPostList 获取帖子列表，分页展示
func GetPostList(page, size int64) (apiPostDetail []*models.ApiPostDetail, err error) {
	// 1. 调用dao层查询数据,到post表里面查询所有数据
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}

	// 初始化apiPostDetail
	apiPostDetail = make([]*models.ApiPostDetail, 0, len(posts))

	// 2. 根据post表的数据，查询其他表，拼接成ApiPostDetail
	// 遍历posts
	for _, post := range posts {
		// 查询authorName(根据post的author_id)
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("AuthorId:", post.AuthorID),
				zap.Error(err))
			continue
		}

		// 查询Community(根据post的community_id)
		communityDetail, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityById(post.CommunityID) failed", zap.Error(err))
			continue
		}
		apiPostDetail = append(apiPostDetail, &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: communityDetail,
		})
	}

	return

}

// GetPostList2 获取帖子列表，根据order决定排序顺序
func GetPostList2(p *models.ParamPostList) (apiPostDetail []*models.ApiPostDetail, err error) {
	// 去redis获取id列表
	ids, err := redis.GetPostIDsInOrder(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder(p) failed", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return nil, err
	}
	zap.L().Debug("GetPostIDsInOrder", zap.Any("ids", ids))
	// 根据ids查询每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	zap.L().Debug("GetPostVoteData", zap.Any("voteData", voteData))
	if err != nil {
		zap.L().Error("redis.GetPostVoteData(ids) failed", zap.Error(err))
		return nil, err
	}

	// 根据id列表去mysql查询post详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids) failed", zap.Error(err))
		return nil, err
	}

	// 初始化apiPostDetail
	apiPostDetail = make([]*models.ApiPostDetail, 0, len(posts))

	// 2. 根据post表的数据，查询其他表，拼接成ApiPostDetail
	// 遍历posts
	for idx, post := range posts {
		// 查询authorName(根据post的author_id)
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("AuthorId:", post.AuthorID),
				zap.Error(err))
			continue
		}

		// 查询Community(根据post的community_id)
		communityDetail, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityById(post.CommunityID) failed", zap.Error(err))
			continue
		}
		apiPostDetail = append(apiPostDetail, &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		})
	}
	return
}

// GetCommunityPostList 按社区id获取帖子列表，根据order决定排序顺序
func GetCommunityPostList(p *models.ParamPostList) (apiPostDetail []*models.ApiPostDetail, err error) {
	// 去redis获取id列表
	ids, err := redis.GetPostIDsInOrderByCommunity(p)
	if err != nil {
		zap.L().Error("redis.GetPostIDsInOrder(p) failed", zap.Error(err))
		return nil, err
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
		return nil, err
	}
	zap.L().Debug("GetPostIDsInOrder", zap.Any("ids", ids))
	// 根据ids查询每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	zap.L().Debug("GetPostVoteData", zap.Any("voteData", voteData))
	if err != nil {
		zap.L().Error("redis.GetPostVoteData(ids) failed", zap.Error(err))
		return nil, err
	}

	// 根据id列表去mysql查询post详细信息
	posts, err := mysql.GetPostListByIDs(ids)
	if err != nil {
		zap.L().Error("mysql.GetPostListByIDs(ids) failed", zap.Error(err))
		return nil, err
	}

	// 初始化apiPostDetail
	apiPostDetail = make([]*models.ApiPostDetail, 0, len(posts))

	// 2. 根据post表的数据，查询其他表，拼接成ApiPostDetail
	// 遍历posts
	for idx, post := range posts {
		// 查询authorName(根据post的author_id)
		user, err := mysql.GetUserById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("AuthorId:", post.AuthorID),
				zap.Error(err))
			continue
		}

		// 查询Community(根据post的community_id)
		communityDetail, err := mysql.GetCommunityById(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityById(post.CommunityID) failed", zap.Error(err))
			continue
		}
		apiPostDetail = append(apiPostDetail, &models.ApiPostDetail{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: communityDetail,
		})
	}
	return
}

// GetPostListNew 根据传入的community参数判断调用哪个处理函数,
// 将GetPostList2和GetCommunityPostList合二为一
func GetPostListNew(p *models.ParamPostList) (apiPostDetail []*models.ApiPostDetail, err error) {
	if p.CommunityID == 0 { // 不含社区id
		apiPostDetail, err = GetPostList2(p)
	} else {
		apiPostDetail, err = GetCommunityPostList(p)
	}
	if err != nil {
		zap.L().Error("GetPostListNew failed", zap.Error(err))
		return nil, err
	}
	return apiPostDetail, err

}
