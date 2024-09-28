package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
	"strconv"
)

// GetBooks 获取数据库中所有的图书
func GetBooks() ([]*model.Book, error) {
	//sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}

		//赋值
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//添加到切片中
		books = append(books, book)
	}

	return books, nil
}

// 向数据库中添加一本图书
func AddBook(b *model.Book) error {
	//sql语句
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImgPath)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书的id从数据库中删除一本图书
func DeleteBook(bookID string) error {
	//sql语句
	sqlStr := "delete from books where id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, bookID)
	if err != nil {
		return err
	}

	return nil
}

// 根据图书的id从数据库中查询出一本图书
func GetBookByID(bookId string) (*model.Book, error) {
	//sql语句
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookId)
	//创建一个book
	book := &model.Book{}
	//字段赋值
	row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)

	return book, nil
}

// 根据图书的id修改图书信息
func UpdateBook(b *model.Book) error {
	//sql语句
	sqlStr := "update books set title=?,author=?,price=?,sales=?,stock=? where id=?"
	//执行
	_, err := utils.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ID)
	if err != nil {
		return nil
	}
	return nil
}

// 获取带分页的图书信息
func GetPageBooks(pageNo string) (*model.Page, error) {
	//类型转换
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总数
	sqlStr := "select count(*) from books"
	//设置变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)

	//设置每页只显示4条记录
	var pageSize int64 = 4

	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页中的图书
	sqlStr2 := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
		//将book添加到books中
		books = append(books, book)
	}

	//创建page实例
	page := &model.Page{
		Books:       books,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
