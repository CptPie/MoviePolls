package common

import (
	"fmt"
	//"time"
	"sort"
)

type Movie struct {
	Id          int
	Name        string
	Links       []string
	Description string

	CycleAdded   *Cycle
	CycleWatched *Cycle

	Removed  bool // Removed by a mod or admin
	Approved bool // Approved by a mod or admin (if required by config)

	Votes []*Vote

	Poster  string // TODO: make this procedural
	AddedBy *User
}

func (m Movie) UserVoted(userId int) bool {
	for _, v := range m.Votes {
		if v.User.Id == userId {
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

	return fmt.Sprintf("Movie{Id:%d Name:%q Links:%s Description:%q CycleAdded:%s CycleWatched:%s Votes:%s}",
		m.Id,
		m.Name,
		m.Links,
		m.Description,
		m.CycleAdded,
		m.CycleWatched,
		votes,
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
