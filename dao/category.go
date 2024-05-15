package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{NewDBClient(ctx)}
}

func NewCategoryDaoByDB(db *gorm.DB) *CategoryDao {
	return &CategoryDao{db}
}

func (categoryDao *CategoryDao) ListCategory() (category []*model.Category, err error) {
	err = categoryDao.DB.Model(&model.Category{}).Find(&category).Error
	return
}
