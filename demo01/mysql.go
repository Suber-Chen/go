package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	d := ConnectMysql()
	// InsertMysql(d)
	// UpdateMysql(d)
	// DeleteMysql(d,1)
	QueryMysqlById(d,2)
	// QueryMysql(d)
}

func ConnectMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(192.168.1.20:3306)/suber")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	// 最大连接时长
	db.SetConnMaxLifetime(time.Minute * 3)
	// 最大连接数
	db.SetMaxOpenConns(10)
	// 空闲连接数
	db.SetMaxIdleConns(10)
	fmt.Println("数据库连接成功...")
	return db
}
