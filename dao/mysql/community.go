package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlstr := "select community_id, community_name from community"
	if err = db.Select(&communityList, sqlstr); err != nil { // community注意要引用！
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no community in db")
			err = nil
		}
		return nil, err
	}
	return
}

func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail) // 为community分配内存
	sqlstr := `select community_id, community_name, introduction, create_time from community where community_id=?`
	if err = db.Get(community, sqlstr, id); err != nil { // db.Get用于查找单个记录，未找到时返回错误，db.Select用于多个记录，未找到时返回空列表
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no community in db")
			err = ErrorInvalidID
		}
	}
	return
}
