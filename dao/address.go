package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao AddressDao) IsOrNotAddress(uid uint, address string) (exist bool, err error) {
	var count int64
	err = dao.DB.Model(&model.Address{}).Where("user_id = ? and address = ?", uid, address).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, err
	}
	return true, err
}

func (dao AddressDao) CreateAddress(address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&address).Error
}

func (dao AddressDao) GetAddressById(addressId, uid uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", addressId, uid).Find(&address).Error
	return address, err
}

func (dao AddressDao) DeleteAddress(id uint) (err error) {
	return dao.DB.Model(&model.Address{}).Where("id = ?", id).Delete(&model.Address{}).Error
}

func (dao AddressDao) UpdateAddress(address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where("id=? AND user_id=?", address.ID, address.UserId).Updates(address).Error
}
