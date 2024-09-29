package controller

import (
	"bookstore001/dao"
	"bookstore001/model"
	"html/template"
	"net/http"
	"strconv"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中带分页的图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, page)
}

//// 获取所有图书
//func GetBooks(w http.ResponseWriter, r *http.Request) {
//	//调用bookdao获取所有图书的函数
//	books, _ := dao.GetBooks()
//	//解析模板文件
//	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
//	//执行
//	t.Execute(w, books)
//}

//// 添加图书
//func AddBook(w http.ResponseWriter, r *http.Request) {
//	//获取图书信息
//	title := r.PostFormValue("title")
//	author := r.PostFormValue("author")
//	price := r.PostFormValue("price")
//	sales := r.PostFormValue("sales")
//	stock := r.PostFormValue("stock")
//	//类型转换
//	fPrice, _ := strconv.ParseFloat(price, 64)
//	iSales, _ := strconv.ParseInt(sales, 10, 0)
//	iStock, _ := strconv.ParseInt(stock, 10, 0)
//	//创建book
//	book := &model.Book{
//		Title:   title,
//		Author:  author,
//		Price:   fPrice,
//		Sales:   int(iSales),
//		Stock:   int(iStock),
//		ImgPath: "/static/img/default.jpg",
//	}
//	//调用bookdao中添加图书的函数
//	dao.AddBook(book)
//	//调用GetBooks函数再次查询一次列表
//	GetBooks(w, r)
//}

// 删除图书
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	//获取要删除的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中删除图书的函数
	dao.DeleteBook(bookID)
	//调用GetBooks函数再次查询一次列表
	GetPageBooks(w, r)
}

// 去添加、修改图书的页面
func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	//获取要修改的图书的id
	bookID := r.FormValue("bookId")
	//调用bookdao中获取图书的函数
	book, _ := dao.GetBookByID(bookID)
	if book.ID > 0 { //有传id
		//更新图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, book)
	} else {
		//添加图书
		//解析模板
		t := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		//执行
		t.Execute(w, "")
	}

}

// 修改或添加图书
func UpdateOrAddBook(w http.ResponseWriter, r *http.Request) {
	//获取图书信息
	bookID := r.PostFormValue("bookId")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	//类型转换
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10, 0)
	iStock, _ := strconv.ParseInt(stock, 10, 0)
	ibookID, _ := strconv.ParseInt(bookID, 10, 0)
	//创建book
	book := &model.Book{
		ID:      int(ibookID),
		Title:   title,
		Author:  author,
		Price:   fPrice,
		Sales:   int(iSales),
		Stock:   int(iStock),
		ImgPath: "/static/img/default.jpg",
	}

	if book.ID > 0 {
		//调用bookdao中更新图书的函数
		dao.UpdateBook(book)
	} else {
		//调用bookdao中添加图书的函数
		dao.AddBook(book)
	}

	//调用GetBooks处理器函数再次查询一次数据库
	GetPageBooks(w, r)
}

// 获取带分页的图书
func GetPageBooks(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	//调用bookdao中带分页的图书的函数
	page, _ := dao.GetPageBooks(pageNo)
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, page)
}
