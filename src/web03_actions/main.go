package main

import (
	"html/template"
	"net/http"
	"time"
	"web03_actions/model"
)

func testIf(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("if.html"))
	age := 17
	//执行
	t.Execute(w, age > 18)
}

func testRange(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("range.html"))
	var emps []*model.Employee

	emp := &model.Employee{
		ID:       1,
		LastName: "宋江",
		Email:    "jsy@ls.com",
	}
	emps = append(emps, emp)
	emp = &model.Employee{
		ID:       2,
		LastName: "卢俊义",
		Email:    "ljy@ls.com",
	}
	emps = append(emps, emp)
	//执行
	t.Execute(w, emps)
}

func testRangeMap(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("rangemap.html"))
	var emps map[string]string
	emps = make(map[string]string)
	emps["神行太保"] = "戴宗"
	emps["浪子"] = "燕青"
	emps["赤发鬼"] = "刘唐"
	//执行
	t.Execute(w, emps)
}

func formatDate(t time.Time) string {
	format := "2006年01月02日"
	return t.Format(format)
}

func actionfunc(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{ // 定义映射
		"chndate": formatDate,
	}
	t := template.New("rangechannel.html").Funcs(funcMap) // 绑定映射到模板
	t, _ = t.ParseFiles("rangechannel.html")
	t.Execute(w, time.Now())
}

func testWith(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("with.html"))
	//执行
	t.Execute(w, "狸猫")
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	t := template.Must(template.ParseFiles("template1.html", "template2.html"))
	//执行
	t.Execute(w, "www")
}

func main() {
	http.HandleFunc("/testIf", testIf)
	http.HandleFunc("/testRange", testRange)

	http.HandleFunc("/testRangeMap", testRangeMap)
	http.HandleFunc("/testRangeChannel", actionfunc)

	http.HandleFunc("/testWith", testWith)
	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
