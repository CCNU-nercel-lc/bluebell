package controller

import (
	"bluebell/dao/mysql"
	"bluebell/settings"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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
	err := mysql.Init(dbCfg)
	if err != nil {
		var t *testing.T
		t.Fatalf("Init failed, err : %v\n", err)
		panic(err)
	}
}

func TestCreatePostHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "http://127.0.0.1:8080/api/v1/createPost"
	r.POST(url, CreatePostHandler)

	body := `{
		"community_id": 1,
		"title":	"test", 
		"content":	"just a test"
	}
	`
	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	//assert.Contains(t, w.Body.String(), "success")

	return

}
