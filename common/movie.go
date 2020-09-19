package common

import (
	"fmt"
	//"time"
	"sort"

	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Name        string
	Links       []*Link `gorm:"many2many:movies_links"`
	Description string
	Remarks     string
	Duration    string
	Rating      float32

	CycleAddedId int
	CycleAdded   *Cycle `gorm:"foreignKey:CycleAddedId"`
	Watched bool // Movie has been watched

	Removed  bool // Removed by a mod or admin
	Approved bool // Approved by a mod or admin (if required by config)

	Votes []*Vote `gorm:"foreignKey:MovieId"`
	Tags  []*Tag  `gorm:"many2many:movies_tags"`

	Poster    string // TODO: make this procedural
	AddedById int
	AddedBy   *User `gorm:"foreignKey:AddedById"`
}

func (m Movie) UserVoted(userId int) bool {
	for _, v := range m.Votes {
		if int(v.User.ID) == userId {
			return true
		}
	}
	return false
}

func (m Movie) String() string {
	votes := []string{}
	for _, v := range m.Votes {
		votes = append(votes, v.User.Name)
	}

	tags := []string{}
	for _, t := range m.Tags {
		tags = append(tags, t.Name)
	}

	return fmt.Sprintf("Movie{Id:%d Name:%q Links:%v Description:%q Remarks:%s CycleAdded:%s Watched:%b Duration:%s Rating:%f Votes:%s Tags:%s}",
		m.ID,
		m.Name,
		m.Links,
		m.Description,
		m.Remarks,
		m.CycleAdded,
		m.Watched,
		m.Duration,
		m.Rating,
		votes,
		tags,
	)
}

type movieVoteSort []*Movie

func (ml movieVoteSort) Len() int      { return len(ml) }
func (ml movieVoteSort) Swap(i, j int) { ml[i], ml[j] = ml[j], ml[i] }

// Sort by votes descending then by name for ties.
func (ml movieVoteSort) Less(i, j int) bool {
	if ml[i].Votes == nil && ml[j].Votes == nil {
		return ml[i].Name < ml[j].Name
	}

	if ml[i].Votes == nil && ml[j].Votes != nil {
		return true
	}

	if ml[i].Votes != nil && ml[j].Votes == nil {
		return false
	}

	if len(ml[i].Votes) == len(ml[j].Votes) {
		return ml[i].Name < ml[j].Name
	}

	return len(ml[i].Votes) > len(ml[j].Votes)
}

func SortMoviesByVotes(list []*Movie) []*Movie {
	s := movieVoteSort(list)
	sort.Sort(s)
	return s
}

type movieNameSort []*Movie

func (ml movieNameSort) Len() int           { return len(ml) }
func (ml movieNameSort) Less(i, j int) bool { return ml[i].Name < ml[j].Name }
func (ml movieNameSort) Swap(i, j int)      { ml[i], ml[j] = ml[j], ml[i] }

func SortMoviesByName(list []*Movie) []*Movie {
	s := movieNameSort(list)
	sort.Sort(s)
	return s
}
