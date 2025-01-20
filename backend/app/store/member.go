package store

import (
	"log/slog"

	"glissando/db"
	"glissando/types"
)

func GetMembers() (types.APIMembers, error) {

	var members types.APIMembers
	var member types.APIMember

	query := db.Connection.Table("members")
	query.Select("members.id, members.name")
	query.Where("members.deleted_at IS NULL")
	query.Order("members.id")
	rows, err := query.Rows()
	if err != nil {
		slog.Error(err.Error())
		return members, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&member.ID, &member.Name); err != nil {
			slog.Error(err.Error())
			return members, err
		}
		members.Items = append(members.Items, member)
		count++
	}
	members.Quantity = count

	return members, nil
}
