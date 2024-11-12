package model

// 分页逻辑
type BasePage struct {
	pageNum  int `form:"pageNum"`
	pageSize int `form:"pageSize"`
}
