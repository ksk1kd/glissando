package seeds

import (
	"glissando/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateProject(db *gorm.DB, name string) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Project{
		Name: name,
	}).Error
}
