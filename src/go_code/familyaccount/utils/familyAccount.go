package utils

import "fmt"

type familyAccount struct {
	//声明必须的字段..

	//声明一个字段，保存接收用户输入的选项
	key string
	//声明一个变量，控制是否退出循环
	loop bool

	//定义账户余额[]
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义一个字段，记录是否有收支的行为
	flag bool

	//当有收支时，只需要对details 进行拼接处理即可
	//收支详情，使用字符串记录
	details string
}

//编写一个工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *familyAccount {
	return &familyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说明",
	}
}

//将显示明细写成一个方法
func (this *familyAccount) showDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支记录...来一笔吧！")
	}
}

//将登记收入写成一个方法，和*FamilyAccount绑定
func (this *familyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money //修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将登记支出写成一个方法，和*FamilyAccount绑定
func (this *familyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//这里需要做一个必要的判断
	if this.balance < this.money {
		fmt.Println("余额不足")
		//break
	}
	this.balance -= this.money //修改账户余额
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	//将这个收入情况，拼接到details变量
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

//将退出系统写成一个方法，和*FamilyAccount绑定
func (this *familyAccount) exit() {
	fmt.Println("确定退出吗?(y/n)")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("请输入正确的选项..")
	}
	if choice == "y" {
		this.loop = false
	}
}

// 给该结构体绑定相应的方法
// 显示主菜单
func (this *familyAccount) MainMenu() {
	for {
		fmt.Println("\n-----------------家庭收支记账软件-----------------")
		fmt.Println("                 1 收支明细")
		fmt.Println("                 2 登记收入")
		fmt.Println("                 3 登记支出")
		fmt.Println("                 4 退出软件")
		fmt.Print("请选择(1-4):")

		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项..")
		}
	}
}
