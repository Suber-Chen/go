package main

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	id   int
	name string
	age  int
}

func QueryMysql(db *sql.DB) {
	queryStr := "select * from user_t1"
	r, err := db.Query(queryStr)
	if err != nil {
		log.Panic("查询出错......")
	}
	for r.Next() {
		var u User
		err2 := r.Scan(&u.id, &u.name, &u.age)
		if err2 != nil {
			log.Panic("查询出错")
		}
		fmt.Printf("&u.id:%v, &u.name:%v, &u.age:%v\n", u.id, u.name, u.age)
	}
}

func QueryMysqlById(db *sql.DB, id int) {
	queryStr := "select * from user_t1 where id=?"
	r := db.QueryRow(queryStr, id)
	var u User
	err := r.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		log.Print("不存在数据...")
		return
	}
	fmt.Printf("id:\t%v\nname:\t%v\nage:\t%v\n", u.id, u.name, u.age)
}
