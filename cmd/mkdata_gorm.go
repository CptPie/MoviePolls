package main

import (
	"fmt"

	"github.com/zorchenhimer/MoviePolls/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host={host} user={user} password={password} dbname={db} port={port}"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error while connecting to DB: %v\n", err)
	}

	err = db.AutoMigrate(&common.User{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&common.Tag{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&common.Movie{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&common.Vote{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&common.Cycle{})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.AutoMigrate(&common.Link{})
	if err != nil {
		fmt.Println(err)
		return
	}

}
