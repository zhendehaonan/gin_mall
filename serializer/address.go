package serializer

import "gin_mall/model"

type Address struct {
	Id        uint   `json:"id"`
	UserId    uint   `json:"user_id_id"`
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
}

func BuildAddress(item *model.Address) Address {
	return Address{
		Id:        item.ID,
		UserId:    item.UserId,
		Name:      item.Name,
		CreatedAt: item.CreatedAt.Unix(),
		Phone:     item.Phone,
		Address:   item.Address,
	}
}
