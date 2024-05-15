package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// 根据id查找notice
func (noticeDao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = noticeDao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
