package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	Id       int64 `xorm:"id"`
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
	// Insert(*engine)
	// Get(*engine)
	// Find(*engine)
	// Count(*engine)
	// Update(*engine)
	// Delete(*engine)
	// Table(*engine)
	// Select(*engine)
	// SetID(*engine)
	// Sum(*engine)
	// SumsInt(*engine)
	// WhereAnd(*engine)
	// WhereOr(*engine)
	// OrderBy(*engine)
	// Asc(*engine)
	// Desc(*engine)
	SQL(*engine)
	Query(*engine)

}

// Insert
func Insert(engine xorm.Engine) {
	user := User{
		Name:     "太郎",
		Password: "パスワード",
		Age:      20,
	}
	_, err := engine.Table("user").Insert(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:", user)
}

//Get 単体取得(1レコードを取得)
func Get(engine xorm.Engine) {
	user := User{}
	result, err := engine.Where("id = ?", 1).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:", user)
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
	result, err := engine.Where("id=?", 1).Delete(&user)
	if err != nil {
		log.Println(err)
	}
	if result == 0 {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", user)
	fmt.Println(result)
}

func Table(engine xorm.Engine) {
	user := User{}
	// ここでテーブルをユーザーテーブルを指定します。
	session := engine.Table("user")
	_, err := session.Where("id = ?", 1).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:", user)
}

func Select(engine xorm.Engine) {
	user := User{}
	// ここでテーブルをユーザーテーブルを指定します。
	session := engine.Table("user")
	result, err := session.Select("name").Where("id = ?", 1).Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user:", user)
	fmt.Println("result:", result)
}

func Sum(engine xorm.Engine) {
	user := User{}
	sumFloat64, _ := engine.Sum(user, "id")
	print(" sumFloat64: ", sumFloat64)
}

func SumsInt(engine xorm.Engine) {
	user := User{}
	sumInt, _ := engine.SumInt(user, "id")
	print(" sumInt: ", sumInt)
}

func WhereAnd(engine xorm.Engine) {
	users := []User{}
	session := engine.Table("user")
	err := session.Where("id > ?", 1).And("age = ?", 30).Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("whereAndUsers:", users)
}

func WhereOr(engine xorm.Engine) {
	users := []User{}
	session := engine.Table("user")
	err := session.Where("id > ?", 1).Or("age = ?", 30).Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("whereOrUsers:", users)
}

func OrderBy(engine xorm.Engine) {
	users := []User{}
	session := engine.Table("user")
	err := session.OrderBy("id desc").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

func Asc(engine xorm.Engine) {
	users := []User{}
	session := engine.Table("user")
	err := session.Asc("id").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("AscUsers:", users)
}

func Desc(engine xorm.Engine) {
	users := []User{}
	session := engine.Table("user")
	err := session.Desc("id").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DescUsers:", users)
}

func SQL(engine xorm.Engine) {
	users := []User{}
	err := engine.SQL("select * from user").Find(&users)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("users:", users)
}

func Query(engine xorm.Engine) {
	results, err := engine.Query("select * from user")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("results:", results)
}
