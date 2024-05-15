package model

//分页

type BasePage struct {
	PageNum  int `form:"page_num"`  //第几页
	PageSize int `form:"page_size"` //每页多少条数据
}
