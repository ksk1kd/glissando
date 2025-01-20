package models

import "gorm.io/gorm"

type Resource struct {
	gorm.Model
	MemberID  uint64 `gorm:"not null"`
	Member    Member
	ProjectID uint64 `gorm:"not null"`
	Project   Project
	Month     string `gorm:"not null"`
	Time      uint64 `gorm:"not null"`
}
