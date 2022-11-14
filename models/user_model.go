package models

import (
	"blog/utils"
	"fmt"
)

// User 用户模型
type User struct {
	Id         int
	Username   string
	Password   string
	Status     int // 0 正常状态， 1删除
	Createtime int64
}

//--------------DB Operation-----------------------

func InsertUser(user User) (int64, error) {
	return utils.ModifyDB("insert into users(username,password,status,createtime) values (?,?,?,?)", user.Username,
		user.Password, user.Status, user.Createtime)
}

// QueryUserWightCon 按条件查询，返回用户id
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

// QueryUserWithUsername 基于用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

// QueryUserWithParam 基于用户名和密码判断登录
func QueryUserWithParam(username string, password string) int {
	sql := fmt.Sprintf("where user='%s' and password='%s'", username, password)
	return QueryUserWightCon(sql)
}
