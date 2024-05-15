package serializer

import (
	"gin_mall/conf"
	"gin_mall/model"
)

type ProductImg struct {
	ProductId uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImgRe(item *model.ProductImg) ProductImg {
	return ProductImg{
		ProductId: item.ProductID,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + item.ImgPath,
	}
}

func BuildProductImgRes(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		product := BuildProductImgRe(item)
		productImgs = append(productImgs, product)
	}
	return productImgs
}
