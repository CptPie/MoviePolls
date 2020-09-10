package common

import (
	"fmt"

	"gorm.io/gorm"
)

type Vote struct {
	gorm.Model

	UserId  int
	User    *User `gorm:"foreignKey:UserId"`
	MovieId int
	Movie   *Movie `gorm:"foreignKey:MovieId"`
	// Decay based on cycles active.
	CycleAddedId int
	CycleAdded   *Cycle `gorm:"foreignKey:CycleAddedId"`
}

func (v Vote) String() string {
	uid := 0
	mid := 0
	cid := 0

	if v.User != nil {
		uid = int(v.User.ID)
	}
	if v.Movie != nil {
		mid = int(v.Movie.ID)
	}
	if v.CycleAdded != nil {
		cid = int(v.CycleAdded.ID)
	}

	return fmt.Sprintf("{Vote User:%d Movie:%d Cycle:%d}", uid, mid, cid)
}
