package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id       int64
	Name     string
	Age      int
	Password string    `xorm:"varchar(200)"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
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

	// Insert
	user := User{
		Name:     "名前",
		Password: "パスワード",
		Age:      20,
	}
	_, err = engine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Insert終了しました。")
	fmt.Println("user:", user)

	//Get 単体取得(1レコードを取得)
	getUser := User{}
	got, err := engine.Where("id = ?", 1).Get(&getUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got:", got)
	fmt.Println("getUser:", getUser)
	if !got {
		log.Fatal("Not Found")
	}

	// Find 複数取得(複数レコードを取得)
	getUsers := []User{}
	err = engine.Where("age = ?", 20).Find(&getUsers)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("err:", err)
	fmt.Println("getUsers:", getUsers)
	if !got {
		log.Fatal("Not Found")
	}

	// Count レコードの数を取得
	countUser := User{}
	count, err := engine.Count(&countUser)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("レコード数:", count)

	// Update
	updateUser := User{
		Name:     "更新名前",
		Password: "更新パスワード",
		Age:      20,
	}
	gots, err := engine.Where("id =?", 1).Update(&updateUser)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("gots:", gots)
	fmt.Println("updateUser:", updateUser)

	// Delete
	deleteUser := User{}
	deleter, err := engine.Where("id=?", 11).Delete(&deleteUser)
	fmt.Println("deleter:", deleter)
	fmt.Println("deleteUser:", deleter)

}
