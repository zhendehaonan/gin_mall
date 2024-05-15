package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

type Product struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	CategoryId    uint   `json:"category_id"`
	Title         string `json:"title"`
	Info          string `json:"info"`
	ImgPath       string `json:"img_path"`
	Price         string `json:"price"`
	DiscountPrice string `json:"discount_price"`
	View          uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
	Num           int    `json:"num"`
	OnSale        bool   `json:"on_sale"`
	BossId        uint   `json:"boss_id"`
	BossName      string `json:"boss_name"`
	BossAvatar    string `json:"boss_avatar"`
}

func BuildProduct(item *model.Product) Product {
	return Product{
		Id:            item.ID,
		Name:          item.Name,
		CategoryId:    item.CategoryID,
		Title:         item.Title,
		Info:          item.Info,
		ImgPath:       conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
		Price:         item.Price,
		DiscountPrice: item.DiscountPrice,
		View:          item.View(),
		CreatedAt:     item.CreatedAt.Unix(),
		OnSale:        item.OnSale,
		Num:           item.Num,
		BossId:        item.BossID,
		BossName:      item.BossName,
		BossAvatar:    conf.Host + conf.HttpPort + conf.AvatarPAth + item.BossAvatar,
	}
}

func BuildProducts(items []model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(&item)
		products = append(products, product)
	}
	return products
}
