package main

import (
	"fmt"
	"time"

	"github.com/zorchenhimer/MoviePolls/common"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=mp password=mp dbname=moviepolls port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error while connecting to DB: %v\n", err)
	}

	err = db.AutoMigrate(&common.User{})
	if err != nil {
		fmt.Printf("User: %v\n", err)
		return
	}
	err = db.AutoMigrate(&common.Tag{})
	if err != nil {
		fmt.Printf("Tag: %v\n", err)
		return
	}
	err = db.AutoMigrate(&common.Movie{})
	if err != nil {
		fmt.Printf("Movie: %v\n", err)
		return
	}
	err = db.AutoMigrate(&common.Vote{})
	if err != nil {
		fmt.Printf("Vote: %v\n", err)
		return
	}
	err = db.AutoMigrate(&common.Cycle{})
	if err != nil {
		fmt.Printf("Cycle: %v\n", err)
		return
	}
	err = db.AutoMigrate(&common.Link{})
	if err != nil {
		fmt.Printf("Link: %v\n", err)
		return
	}

	tag1 := common.Tag{Name: "foo"}

	tag2 := common.Tag{Name: "bar"}

	// 	db.Create(&tag1)
	// 	db.Create(&tag2)

	tags := []*common.Tag{&tag1, &tag2}

	link1 := common.Link{Url: "imdb.com", Type: "imdb", IsSource: true}
	link2 := common.Link{Url: "foo-bar.com", Type: "misc", IsSource: false}

	// 	db.Create(&link1)
	// 	db.Create(&link2)

	links := []*common.Link{&link1, &link2}

	now := time.Now()

	cycleadded := common.Cycle{
		PlannedEnd: &now,
	}

	// db.Create(&cycleadded)

	user := common.User{
		Name:                "testuser",
		Password:            "lkjklj;kljl;",
		OAuthToken:          "",
		Email:               "",
		NotifyCycleEnd:      false,
		NotifyVoteSelection: false,
		Privilege:           0,
		PassDate:            now,
		LastMovieAdd:        now,
	}

	// db.Create(&user)

	// cyclewatched := common.Cycle{
	// 	PlannedEnd: &now,
	// 	Ended: &now,
	// }

	movie := common.Movie{
		Name:        "testmovie",
		Links:       links,
		Description: "example Description",
		Remarks:     "some fancy remarks",
		Duration:    "1h30m",
		Rating:      7.8,
		CycleAdded:  &cycleadded,
		Removed:     false,
		Approved:    false,
		Tags:        tags,
		Poster:      "/foo.bar",
		AddedBy:     &user,
	}

	id := db.Create(&movie)

	var found_movie common.Movie
	db.Preload("Tags").Preload("Links").Preload("AddedBy").Find(&found_movie, id)
	fmt.Printf("%v\n", found_movie)

}
