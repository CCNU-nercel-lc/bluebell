package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"testing"
)

func init() {
	dbCfg := &settings.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "root",
		DbName:       "bluebell",
		Port:         3306,
		MaxOpenConns: 200,
		MaxIdleConns: 50,
	}
	err := Init(dbCfg)
	if err != nil {
		var t *testing.T
		t.Fatalf("Init failed, err : %v\n", err)
		panic(err)
	}
}

func TestCreatePost(t *testing.T) {
	post := models.Post{
		Title:       "test",
		Content:     "test",
		ID:          10,
		AuthorID:    123,
		CommunityID: 1,
	}
	err := CreatePost(&post)
	if err != nil {
		t.Fatalf("CreatePost failed, err : %v\n", err)
	}
}
