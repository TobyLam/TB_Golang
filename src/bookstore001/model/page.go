package model

// 分页结构体
type Page struct {
	Books       []*Book //每页查询出来的图书存放数据切片
	PageNo      int64   //当前页码
	PageSize    int64   //每页条数
	TotalPageNo int64   //总页数
	TotalRecord int64   //总记录数
}
