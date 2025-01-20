package store

import (
	"log/slog"

	"glissando/db"
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
