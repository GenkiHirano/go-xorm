package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:root@tcp([127.0.0.1]:3306)/sample_db?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

	var user User
	user.Name = "名前"
	user.Passwd = "パスワード"
	engine.Table("user").Insert(user)

	fmt.Println("Insert終了しました。")
}
