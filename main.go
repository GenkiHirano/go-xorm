package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	ID       int64
	Name     string
	Age      int
	Password string    `xorm:"password"`
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

	// メソッドを実行します
	Insert(*engine)
	Get(*engine)
	Find(*engine)
	Count(*engine)
	Update(*engine)
	Delete(*engine)
}

// Insert
func Insert(engine xorm.Engine) {
	newUser := User{
		Name:     "太郎",
		Password: "パスワード",
		Age:      20,
	}
	_, err := engine.Table("user").Insert(newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("newUser:", newUser)
}

//Get 単体取得(1レコードを取得)
func Get(engine xorm.Engine) {
	getUser := User{}
	result, err := engine.Where("id = ?", 1).Get(&getUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("getUser:", getUser)
	if !result {
		log.Fatal("Not Found")
	}
}

// Find 複数取得(複数レコードを取得)
func Find(engine xorm.Engine) {
	users := []User{}
	// ageが20のuserを全件取得します
	err := engine.Where("age = ?", 20).Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

// Count レコードの数を取得
func Count(engine xorm.Engine) {
	user := User{}
	count, err := engine.Count(&user)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("レコード数:", count)
}

// Update
func Update(engine xorm.Engine) {
	user := User{
		Name:     "更新した名前",
		Password: "更新したパスワード",
		Age:      30,
	}
	result, err := engine.Where("id =?", 1).Update(&user)
	if err != nil {
		log.Println(err)
	}
	if result == 0 {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", user)
}

// Delete
func Delete(engine xorm.Engine) {
	user := User{}
	result, err := engine.Where("id=?", 11).Delete(&user)
	if err != nil {
		log.Println(err)
	}
	if result == 0 {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", user)
}
