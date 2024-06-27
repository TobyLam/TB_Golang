package model
import(
	"fmt"
)

//定义一个结构体account
type account struct{
	accountNo string
	pwd string
	balance float64
}

//工厂模式的函数-构造函数
func NewAccount(accountNo string,pwd string,balance float64) *account {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return nil
	}
	 
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额数目不对...")
		return nil
	}

	return &account{
		accountNo : accountNo,
		pwd : pwd,
		balance : balance,
	}
}

func (a *account) SetPwd(pwd string) {
	if len(pwd) != 6 {
		fmt.Println("密码的长度不对...")
		return
	}
	a.pwd = pwd
}
func (a *account) GetPwd() string {
	return a.pwd
}

func (a *account) SetBalance(balance float64) {
	if balance < 20 {
		fmt.Println("余额数目不对...")
		return
	}
	a.balance = balance
}
func (a *account) GetBalance() float64 {
	return a.balance
}

func (a *account) SetAccountNo(accountNo string) {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号的长度不对...")
		return
	}
	a.accountNo = accountNo
}
func (a *account) GetAccountNo() string {
	return a.accountNo
}

//方法
//1.存款
func (account *account) Deposite(money float64,pwd string) {
	//看看输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看存款金额是否正确
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.balance += money
	fmt.Println("存款成功")
}

//2.取款
func (account *account) Withdraw(money float64,pwd string) {
	//看看输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	//看看取款金额是否正确
	if money <= 0 || money > account.balance {
		fmt.Println("你输入的金额不正确")
		return
	}

	account.balance -= money
	fmt.Println("取款成功")
}

//3.查询余额
func (account *account) Query(pwd string) {
	//看看输入的密码是否正确
	if pwd != account.pwd {
		fmt.Println("你输入的密码不正确")
		return
	}

	fmt.Printf("你的账号为%v,余额为%v\n",account.accountNo,account.balance)
}