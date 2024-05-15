package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type AddressService struct {
	Address string `json:"address" form:"address"`
	Phone   string `json:"phone" form:"phone"`
	Name    string `json:"name" form:"name"`
	model.BasePage
}

// 创建地址
func (service *AddressService) Create(ctx context.Context, uid uint) serializer.Response {
	var err error
	code := e.Success
	//判断地址 手机号 用户名不能为空
	if service.Address == "" || service.Phone == "" || service.Name == "" || len(service.Phone) != 11 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	//判断该地址是否已经存在
	exist, err := addressDao.IsOrNotAddress(uid, service.Address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    e.GetMsg(code),
		}
	}
	//存在
	if exist {
		code = e.ErrorAddressExist
		return serializer.Response{
			Status: code,
			Error:  e.GetMsg(code),
		}
	}
	//不存在
	address := &model.Address{
		UserId:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err = addressDao.CreateAddress(address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Error:  err.Error(),
			Msg:    "创建失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "创建成功",
	}
}

// 删除地址
func (service *AddressService) Delete(ctx context.Context, uid uint, aid string) serializer.Response {
	var err error
	code := e.Success
	addressId, _ := strconv.ParseInt(aid, 10, 64)
	addressDao := dao.NewAddressDao(ctx)
	//判断该地址是否存在
	_, err = addressDao.GetAddressById(uint(addressId), uid)
	//不存在
	if err != nil {
		code = e.ErrorAddressNotExist
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//存在
	err = addressDao.DeleteAddress(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    "删除失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "删除成功",
	}
}

// 修改地址
func (service *AddressService) Update(ctx context.Context, uid uint, aId string) serializer.Response {
	var err error
	code := e.Success
	addressId, _ := strconv.ParseInt(aId, 10, 64)
	//判断地址 手机号 用户名 不能为空 手机号长度为11
	if service.Address == "" || service.Phone == "" || service.Name == "" || len(service.Phone) != 11 {
		code = e.InvalidParams
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(uint(addressId), uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	address.Address = service.Address
	address.Phone = service.Phone
	address.Name = service.Name
	err = addressDao.UpdateAddress(address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    "修改成功",
	}
}

func (service *AddressService) Show(ctx context.Context, uid uint, aId string) serializer.Response {
	var err error
	code := e.Success
	addressId, _ := strconv.ParseInt(aId, 10, 64)
	addressDao := dao.NewAddressDao(ctx)
	address, err := addressDao.GetAddressById(uint(addressId), uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildAddress(address),
	}
}
