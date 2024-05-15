package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/pkg/util"
	"gin_mall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	Id            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryId    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

// 创建商品和商品图片
func (service *ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserById(uid)
	//以第一张图作为封面图
	tmp, err := files[0].Open()
	path, err := UploadProductToLocalStatic(tmp, uid, service.Name) //上传第一张图片到本地
	if err != nil {
		code = e.ErrorProductImgUpload
		util.LogrusObj.Infoln("product CreateProduct service err:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryId,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossID:        uid,
		BossName:      boss.UserName,
		BossAvatar:    boss.Avatar,
	}
	//创建商品
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.ErrorProductCreate
		util.LogrusObj.Infoln("product CreateProduct service err:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//创建商品图片(因为不止一张图片，所以采用并发)
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for index, file := range files {
		num := strconv.Itoa(index)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ := file.Open()
		path, err = UploadProductToLocalStatic(tmp, uid, service.Name+num) //上传所有图片到本地
		if err != nil {
			code = e.ErrorProductImgUpload
			util.LogrusObj.Infoln("product CreateProduct service err:", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.Error
			util.LogrusObj.Infoln("product CreateProduct service err:", err)
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(product),
	}
}

// 对商品分页处理
func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []model.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryId != 0 {
		condition["category_id"] = service.CategoryId
	}
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.GetCountProductByCondition(condition) //获取商品总数
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("product list service err:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao := dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.GetListProductByCondition(condition, service.BasePage) //获取商品
		wg.Done()
	}()
	wg.Wait()
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

// 搜素商品
func (service *ProductService) Search(ctx context.Context) serializer.Response {
	var product model.Product
	var err error
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	if service.Name == "" {
		code = e.Error
		util.LogrusObj.Infoln("product search service err:", err)
		return serializer.Response{
			Status: code,
			Msg:    "商品名不能为空",
		}
	}
	product, err = productDao.GetProductByName(service.Name)
	if product.ID == 0 {
		return serializer.Response{
			Status: e.Error,
			Msg:    "商品不存在",
		}
	}
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("product search service err:", err)
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(&product),
	}
}

// 商品详情
func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	var product model.Product
	var err error
	code := e.Success
	uid, _ := strconv.Atoi(id)
	productDao := dao.NewProductDao(ctx)
	product, err = productDao.GetProductById(uint(uid))
	if err != nil {
		code = e.Error
		util.LogrusObj.Infoln("product show service err:", err)
		return serializer.Response{
			Status: code,
			Msg:    "商品不存在",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   serializer.BuildProduct(&product),
	}
}
