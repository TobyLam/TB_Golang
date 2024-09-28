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
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	if err != nil {
		return nil, err
	}
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
