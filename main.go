package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"xorm.io/xorm"
)

type User struct {
	Id       int64 `xorm:"id"`
	Name     string
	Age      int
	Password string `xorm:"password"`
}

func main() {

	err := godotenv.Load(".env")
	addr := getAddr(os.Getenv("DATABASE_WRITE_HOST"), os.Getenv("DATABASE_READ_HOST"))

	conf := &mysql.Config{
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               os.Getenv("DATABASE_NAME"),
		User:                 os.Getenv("MAINTENANCE_DATABASE_USER"),
		Passwd:               os.Getenv("MAINTENANCE_DATABASE_PASSWORD"),
		AllowNativePasswords: true,
	}
	engine, err := xorm.NewEngine("mysql", conf.FormatDSN())

	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

	// メソッドを実行します
	// Insert(*engine)
	// Get(*engine)
	Find(*engine)
	// Count(*engine)
	// Update(*engine)
	// Delete(*engine)
}

func getAddr(host, port string) string {
	var addr string
	if strings.Contains(host, ":") {
		addr = host
	} else {
		pInt := 3306
		if p, err := strconv.Atoi(port); err == nil {
			pInt = p
		}
		addr = fmt.Sprintf("%v:%v", host, pInt)
	}
	return addr
}

// Insert
func Insert(engine xorm.Engine) {
	user := User{
		Id:       1,
		Name:     "太郎",
		Age:      20,
		Password: "パスワード",
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
	if !result {
		log.Fatal("Not Found")
	}
	fmt.Println("user:", user)
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
		Id:       1,
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
