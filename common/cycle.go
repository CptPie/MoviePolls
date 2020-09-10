package common

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Cycle struct {
	gorm.Model

	PlannedEnd *time.Time
	Ended      *time.Time

	// List of movies watched this cycle.  If cycle has not ended, this will be
	// nil.
	Watched []*Movie `gorm:"many2many:movies_cycles"`
}

func (c Cycle) PlannedEndString() string {
	if c.PlannedEnd == nil {
		return ""
	}
	return c.PlannedEnd.Format("Mon Jan 2")
}

func (c Cycle) EndedString() string {
	if c.Ended == nil {
		return ""
	}
	return c.Ended.Format("Mon Jan 2, 2006")
}

func (c Cycle) String() string {
	return fmt.Sprintf("Cycle{Id:%d PlannedEnd:%s Ended: %s}", c.ID, c.PlannedEndString(), c.EndedString())
}
