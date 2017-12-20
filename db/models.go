package main

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	err      error
	affected int64
	engine   *xorm.Engine
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
	engine, err = xorm.NewEngine("mysql", "root:123456@(127.0.0.1:3306)/opsweb?charset=utf8")
	if err != nil {
		fmt.Errorf("models.init(fail to conntect database): %v", err)
	}
	if err = engine.Ping(); err != nil {
		log.Fatal(err)
	}
	err = engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}

	/*
		// Insert
		user1 := User{
			Name:    "monkey",
			Salt:    "100",
			Age:     28,
			Passwd:  "123456",
			Created: time.Now(),
			Updated: time.Now(),
		}

		affected, err = engine.Insert(&user1)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(affected)

		var user User
		has, err := engine.Get(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(has)
		fmt.Println(user)
	*/

	// where
	var user User
	has, err := engine.Where("name = ?", "monkey2").Desc("id").Get(&user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(has)
	fmt.Println(user)
}
