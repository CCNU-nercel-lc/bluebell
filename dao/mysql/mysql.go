package mysql

import (
	"bluebell/models"
	"bluebell/settings"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	//user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_cons"))
	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_cons"))
	return
}

func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	// db 查询数据库
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	// 用户存在
	if count > 0 {
		return ErrorUserExist
	}
	// 用户不存在返回空err
	return

}

// InsertUser 添加用户
func InsertUser(user *models.User) (err error) {
	// 对面进行加密
	user.Password = encryptPassword(user.Password)
	// 执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

// 加密密码
func encryptPassword(opassword string) string {
	h := md5.New()
	h.Write([]byte("lc@123.com"))
	return hex.EncodeToString(h.Sum([]byte(opassword)))
}

func Close() {
	_ = db.Close()
}

func Login(user *models.User) (err error) {
	oPassword := user.Password // 保存源密码
	sqlStr := `select user_id, username, password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNotExist
	}
	if err != nil {
		return err // 查询数据库失败
	}
	// 查询成功, 判断密码是否正确
	if user.Password != encryptPassword(oPassword) {
		return ErrorInvalidPassword
	}
	return err

}
