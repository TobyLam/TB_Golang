package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
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
