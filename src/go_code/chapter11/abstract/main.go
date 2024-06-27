package main
import(
	"fmt"
)

//定义一个结构体Account
type Account struct{
	AccountNo string
	Pwd string
	Balance float64
}

//方法
//1.存款
func (account *Account) Deposite(money float64,pwd string) {
	//看看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance += money
	fmt.Println("存款成功")
}

//2.取款
func (account *Account) Withdraw(money float64,pwd string) {
	//看看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看取款金额是否正确
	if money <= 0 || money > account.Balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.Balance -= money
	fmt.Println("取款成功")
}

//3.查询余额
func (account *Account) Query(pwd string) {
	//看看输入的密码是否正确
	if pwd != account.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为%v,余额为%v\n",account.AccountNo,account.Balance)
}

//4.操作
func (account *Account) Handle(options int,pwd string)  {
	switch options {
	case 1:
		account.Deposite(100,pwd)
	case 2:
		account.Withdraw(100,pwd)
	case 3:
		account.Query(pwd)
	case 4:
		fmt.Println("你退出了系统")
	default:
		fmt.Println("你输入的选项不正确")
	}
}

func main(){

	//测试
	account := Account{
		AccountNo:"gs11111111",
		Pwd:"666666",
		Balance:100,
	}
	for {
		//让用户通过控制台输入命令...
		//菜单...
		fmt.Println("1.存款 2.取款 3.查询余额 4.退出")
		var options int
		fmt.Scanln(&options)
		account.Handle(options,account.Pwd)

	}
	
}