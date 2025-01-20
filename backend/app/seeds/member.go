package seeds

import (
	"glissando/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateMember(db *gorm.DB, name string) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Member{
		Name: name,
	}).Error
}
