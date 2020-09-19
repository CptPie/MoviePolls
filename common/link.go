package common

import (
	"fmt"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	IsSource bool
	Type     string
	Url      string
}

func (l Link) String() string {
	return fmt.Sprintf("Link{Url:%s Type:%s IsSource:%v}", l.Url, l.Type, l.IsSource)
}
