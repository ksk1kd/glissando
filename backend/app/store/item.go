package store

import (
	"log/slog"

	"glissando/db"
	"glissando/types"
)

func GetItems(iids []string, q string, category string, creator_id string) (types.APIItems, error) {

	var items types.APIItems
	var item types.APIItem

	query := db.Connection.Table("items")
	query.Select("items.id, items.title, items.image, categories.id, categories.name, items.price, items.creator_id")
	query.Joins("LEFT JOIN categories ON categories.id = items.category_id")
	query.Where("items.deleted_at IS NULL")
	if len(iids) != 0 {
		query.Where("items.id IN ?", iids)
	}
	if q != "" {
		query.Where("items.title LIKE ?", "%"+q+"%")
	}
	if category != "" {
		query.Where("items.category_id = ?", category)
	}
	if creator_id != "" {
		query.Where("items.creator_id = ?", creator_id)
	}
	query.Order("items.id")
	rows, err := query.Rows()
	if err != nil {
		slog.Error(err.Error())
		return items, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&item.ID, &item.Title, &item.Image, &item.Category, &item.CategoryName, &item.Price, &item.Creator); err != nil {
			slog.Error(err.Error())
			return items, err
		}
		items.Items = append(items.Items, item)
		count++
	}
	items.Quantity = count

	return items, nil
}

func SearchItems(q string, category string) (types.APIItems, error) {

	items, err := GetItems(nil, q, category, "")
	if err != nil {
		return items, err
	}

	return items, nil
}
