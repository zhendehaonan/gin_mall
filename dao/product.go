package dao

import (
	"context"
	"gin_mall/model"
	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{NewDBClient(ctx)}
}

func NewProductDaoByDB(db *gorm.DB) *ProductDao {
	return &ProductDao{db}
}

// 创建商品
func (productDao *ProductDao) CreateProduct(product *model.Product) error {
	return productDao.DB.Model(&model.Product{}).Create(&product).Error
}

// 商品总数
func (productDao *ProductDao) GetCountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = productDao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

// 分页
func (productDao *ProductDao) GetListProductByCondition(condition map[string]interface{}, page model.BasePage) (products []model.Product, err error) {
	err = productDao.DB.Where(condition).Offset((page.PageNum - 1) * (page.PageSize)).Limit(page.PageSize).Find(&products).Error
	return
}

// 通过name搜索商品
func (productDao *ProductDao) GetProductByName(name string) (product model.Product, err error) {
	err = productDao.DB.Where("name LIKE ?", "%"+name+"%").First(&product).Error
	return
}

// 通过id搜索商品
func (productDao *ProductDao) GetProductById(id uint) (product model.Product, err error) {
	err = productDao.DB.Where("id = ?", id).First(&product).Error
	return
}
