package serializer

import (
	"gin_mall/model"
)

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreatedAt    int64  `json:"created_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		CreatedAt:    item.CreatedAt.Unix(),
	}
}

func BuildCategories(items []*model.Category) (categories []Category) {
	for _, item := range items {
		category := BuildCategory(item)
		categories = append(categories, category)
	}
	return categories
}
