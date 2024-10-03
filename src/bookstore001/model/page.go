package model

// 分页结构体
type Page struct {
	Books       []*Book //每页查询出来的图书存放数据切片
	PageNo      int64   //当前页码
	PageSize    int64   //每页条数
	TotalPageNo int64   //总页数
	TotalRecord int64   //总记录数
	MinPrice    string
	MaxPrice    string
	IsLogin     bool
	Username    string
}

// 判断是否有上一页
func (this *Page) IsHasPrev() bool {
	return this.PageNo > 1
}

// 判断是否有下一页
func (this *Page) IsHasNext() bool {
	return this.PageNo < this.TotalPageNo
}

// 获取上一页
func (this *Page) GetPrevPageNo() int64 {
	if this.IsHasPrev() {
		return this.PageNo - 1
	} else {
		return 1
	}
}

// 获取下一页
func (this *Page) GetNextPageNo() int64 {
	if this.IsHasNext() {
		return this.PageNo + 1
	} else {
		return this.TotalPageNo
	}
}
