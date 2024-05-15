package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{NewDBClient(ctx)}
}

func NewFavoriteDaoByDB(db *gorm.DB) *FavoriteDao {
	return &FavoriteDao{db}
}

func (dao *FavoriteDao) ListFavorite(id uint) (resp []*model.Favorite, err error) {
	err = dao.DB.Model(&model.Favorite{}).Where("user_id = ?", id).Find(&resp).Error
	return resp, err
}

func (dao *FavoriteDao) IsOrNotExist(pid, uid uint) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Favorite{}).Where("product_id = ? AND user_id = ?", pid, uid).Find(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, err
}

func (dao *FavoriteDao) CreateFavorite(favorite *model.Favorite) error {
	return dao.DB.Model(&model.Favorite{}).Create(&favorite).Error
}

func (dao *FavoriteDao) DeleteFavorite(uid, fId uint) error {
	return dao.DB.Model(&model.Favorite{}).Where("id = ? AND user_id = ? ", fId, uid).Delete(&model.Favorite{}).Error
}
