package common

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	IsSource bool
	Type     string
	Url      string
}
