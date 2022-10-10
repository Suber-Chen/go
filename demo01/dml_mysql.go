package main

import (
	"database/sql"
	"fmt"
	"log"
)

// 插入数据
func InsertMysql(db *sql.DB) {
	queryStr := "insert into user_t1 (name,age) values (?,?)"
	r, err := db.Exec(queryStr, "张三", "33")
	if err != nil {
		log.Panic("插入数据出错")
	}
	i, err2 := r.RowsAffected()
	if err != nil {
		log.Printf("%v", err2.Error())
	}
	fmt.Printf("i: %v行受影响\n", i)
}

func UpdateMysql(db *sql.DB) {
	queryStr := "update user_t1 set name=?,age=? where id=?"
	r, err := db.Exec(queryStr, "suber", 24, 2)
	if err != nil {
		log.Panic("更新失败")
	}
	i, err2 := r.RowsAffected()
	if err != nil {
		log.Printf("%v", err2.Error())
	}
	fmt.Printf("%v行受影响\n", i)
}

func DeleteMysql(db *sql.DB,id int) {
	queryStr := "delete from user_t1 where id = ?"
	r, err := db.Exec(queryStr, id)
	if err != nil {
		log.Panic("删除失败")
	}
	i, err2 := r.RowsAffected()
	if err != nil {
		log.Printf("%v", err2.Error())
	}
	fmt.Printf("i: %v行受影响\n", i)
}
