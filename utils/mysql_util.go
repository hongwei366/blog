package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/prometheus/common/log"

	beego "github.com/beego/beego/v2/server/web"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitMysql() {
	fmt.Println("初始化mysql。。。")
	driverName, _ := beego.AppConfig.String("driverName")
	//数据库连接
	user, _ := beego.AppConfig.String("mysqluser")
	pwd, _ := beego.AppConfig.String("mysqlpwd")
	host, _ := beego.AppConfig.String("host")
	port, _ := beego.AppConfig.String("port")
	dbname, _ := beego.AppConfig.String("dbname")
	//dbConn := "root:yu271400@tcp(127.0.0.1:3306)/myblog?charset=utf8"
	dbConn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	openDb, err := sql.Open(driverName, dbConn)
	if err != nil {
		fmt.Println(err)
	} else {
		db = openDb
		CreateTableWithUser()
	}
}

// CreateTableWithUser 创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
		id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		username VARCHAR(64),
		password VARCHAR(64),
		status INT(4),
		createtime INT(10)
		);`
	ModifyDB(sql)
	log.Info("init table users")
}

// ModifyDB 操作数据库
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	res, err := db.Exec(sql, args...)
	if err != nil {
		fmt.Println(err)
		return 0, err
	} else {
		count, _ := res.RowsAffected()
		return count, nil
	}
}

// QueryRowDB 查询
func QueryRowDB(sql string) *sql.Row {
	return db.QueryRow(sql)
}

// Md5 生成md5值
func Md5(str string) string {
	md := md5.New()
	_, _ = io.WriteString(md, str)
	return hex.EncodeToString(md.Sum(nil))
}
