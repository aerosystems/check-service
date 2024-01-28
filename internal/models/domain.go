package models

import (
	"time"
)

type Domain struct {
	Id        int       `json:"-" gorm:"primaryKey;unique;autoIncrement"`
	Name      string    `json:"name" gorm:"uniqueIndex:idx_name_coverage" example:"gmail.com"`
	Type      string    `json:"type" example:"whitelist"`
	Coverage  string    `json:"coverage" gorm:"uniqueIndex:idx_name_coverage" example:"equals"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
