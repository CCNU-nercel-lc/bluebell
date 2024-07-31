package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查询数据库中的community数据，并返回
	return mysql.GetCommunityList()
}
func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {

	// 查询数据库中给定id的community数据，并返回
	return mysql.GetCommunityDetailByID(id)
}
