package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"), os.Getenv("PGUSER"), os.Getenv("PGDB"), os.Getenv("PGPASS")))

	if err != nil {
		panic(err)
	}

	if !Db.HasTable(&Todo{}) {
		Db.AutoMigrate(&Todo{})
		Db.Create(&Todo{Title: "Todo One", Completed: false})
		Db.Create(&Todo{Title: "Todo Two", Completed: true})
		Db.Create(&Todo{Title: "Todo Three", Completed: false})
	}
}
