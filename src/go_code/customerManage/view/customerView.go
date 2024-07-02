package main

import (
	"fmt"
	"go_code/customerManage/model"
	"go_code/customerManage/service"
)

type customerView struct {

	//定义必要字段
	key  string //接收用户输入...
	loop bool   //表示是否循环的显示主菜单

	//增加一个字段customerService
	//customerService *service.CustomerService
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) list() {
	//首先，获取到当前所有的客户信息（在切片中）
	customers := this.customerService.List()
	//显示
	fmt.Println("-----------------客户列表-----------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n-----------------客户列表完成-----------------\n\n")
}

// 得到用户的输入，构建新的客户信息，并完成添加
func (this *customerView) add() {
	fmt.Println("---------------添加客户----------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Print("邮箱：")
	email := ""
	fmt.Scanln(&email)
	//构造一个新的Customer实例
	//注意，id号没有让用户输入，id是唯一的，需要系统分配
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Add(customer) {
		fmt.Println("------------------------添加完成！--------------------------")
	} else {
		fmt.Println("------------------------添加失败！--------------------------")
	}
}

// 得到用户的输入id，删除该id对应的客户信息
func (this *customerView) delete() {
	fmt.Println("---------------删除客户----------------")
	fmt.Print("请选择待删除客户编号(-1退出)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃删除操作
	}
	fmt.Println("确认是否删除(Y/N)：")
	//可以加入循环，知道输入Y或者N
	choice := ""
	fmt.Scanln(&choice)
	if choice == "Y" || choice == "y" {
		//调用customerService的Delete方法
		if this.customerService.Delete(id) {
			fmt.Println("------------------------删除完成！--------------------------")
		} else {
			fmt.Println("-----------------删除失败！输入的id不存在--------------------")
		}
	} else {
		fmt.Println("------------------------删除取消！--------------------------")
	}
}

// 得到用户的输入id，修改该id对应的客户信息
func (this *customerView) update() {
	fmt.Println("---------------修改客户----------------")
	fmt.Println("请选择待修改客户编号(-1退出,按回车表示不修改当前信息)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return //放弃修改操作
	}
	index := this.customerService.FindById(id)
	customers := this.customerService.List()

	fmt.Printf("姓名(%v)：", customers[index].Name)
	name := ""
	fmt.Scanln(&name)
	if name == "" {
		name = customers[index].Name
	}
	fmt.Printf("性别(%v)：", customers[index].Gender)
	gender := ""
	fmt.Scanln(&gender)
	if gender == "" {
		gender = customers[index].Gender
	}

	fmt.Printf("年龄(%v)：", customers[index].Age)
	age := 0
	fmt.Scanln(&age)
	if age == 0 {
		age = customers[index].Age
	}
	fmt.Printf("电话(%v)：", customers[index].Phone)
	phone := ""
	fmt.Scanln(&phone)
	if phone == "" {
		phone = customers[index].Phone
	}

	fmt.Printf("邮箱(%v)：", customers[index].Email)
	email := ""
	fmt.Scanln(&email)
	if email == "" {
		email = customers[index].Email
	}

	//构造一个新的Customer实例
	customer := model.NewCustomer2(name, gender, age, phone, email)
	//调用
	if this.customerService.Update(customers[index].Id, customer) {
		fmt.Println("------------------------修改完成！--------------------------")
	} else {
		fmt.Println("------------------------修改失败！输入的id不存在--------------------")
	}

}

// 退出软件
func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N)：")
	for {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		} else {
			fmt.Println("选择错误，请重新选择")
		}
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("-----------客户信息管理软件-----------")
		fmt.Println("           1 添加客户")
		fmt.Println("           2 修改客户")
		fmt.Println("           3 删除客户")
		fmt.Println("           4 客户列表")
		fmt.Println("           5 退    出")
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("选择错误，请重新选择")
		}

		if !this.loop {
			break
		}
	}

	fmt.Println("退出客户关系管理系统")
}

func main() {
	//在main函数中，创建一个customerView实例，并运行显示主菜单
	customerView := customerView{
		key:  "",
		loop: true,
	}
	//完成对customerView结构体的customerService字段初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}
