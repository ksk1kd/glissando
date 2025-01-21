package store

import (
	"log/slog"
	"strconv"

	"glissando/db"
	"glissando/models"
	"glissando/types"
)

func GetResources(member string, project string, month string) (types.APIResources, error) {

	var resources types.APIResources
	var resource types.APIResource

	query := db.Connection.Table("resources")
	query.Select("resources.id, resources.time")
	query.Where("resources.deleted_at IS NULL")
	query.Where("resources.member_id = ?", member)
	query.Where("resources.project_id = ?", project)
	query.Where("resources.month = ?", month)
	rows, err := query.Rows()
	if err != nil {
		slog.Error(err.Error())
		return resources, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&resource.ID, &resource.Time); err != nil {
			slog.Error(err.Error())
			return resources, err
		}
		resources.Items = append(resources.Items, resource)
		count++
	}
	resources.Quantity = count

	if count == 0 {
		resources.Items = []types.APIResource{}
	}

	return resources, nil
}

func AddResource(member string, project string, month string, time string) error {

	memberID, _ := strconv.ParseUint(member, 10, 64)
	projectID, _ := strconv.ParseUint(project, 10, 64)
	timeValue, _ := strconv.ParseUint(time, 10, 64)

	resource := models.Resource{
		MemberID:  memberID,
		ProjectID: projectID,
		Month:     month,
		Time:      timeValue,
	}

	if result := db.Connection.Create(&resource); result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}

func UpdateResourceTime(member string, project string, month string, time string) error {

	resources, err := GetResources(member, project, month)
	if err != nil {
		return err
	}

	if len(resources.Items) == 0 {
		return AddResource(member, project, month, time)
	}

	if result := db.Connection.Model(&models.Resource{}).Where("member_id = ?", member).Where("project_id = ?", project).Where("month = ?", month).Update("time", time); result.Error != nil {
		slog.Error(result.Error.Error())
		return result.Error
	}

	return nil
}
