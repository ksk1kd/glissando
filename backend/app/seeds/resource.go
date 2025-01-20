package seeds

import (
	"glissando/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateResource(db *gorm.DB, member_id uint64, project_id uint64, month string, time uint64) error {
	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&models.Resource{
		MemberID:  member_id,
		ProjectID: project_id,
		Month:     month,
		Time:      time,
	}).Error
}
