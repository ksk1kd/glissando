package store

import (
	"log/slog"

	"glissando/db"
	"glissando/types"
)

func GetProjects() (types.APIProjects, error) {

	var projects types.APIProjects
	var project types.APIProject

	query := db.Connection.Table("projects")
	query.Select("projects.id, projects.name")
	query.Where("projects.deleted_at IS NULL")
	query.Order("projects.id")
	rows, err := query.Rows()
	if err != nil {
		slog.Error(err.Error())
		return projects, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&project.ID, &project.Name); err != nil {
			slog.Error(err.Error())
			return projects, err
		}
		projects.Items = append(projects.Items, project)
		count++
	}
	projects.Quantity = count

	return projects, nil
}
