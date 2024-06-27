package main
import(
	"fmt"
	"go_code/chapter11/encapexercise/model"
)

func main(){
	//创建一个account变量
	account := model.NewAccount("jzh11111","000000",40)
	if account != nil{
		fmt.Println("创建成功=",account)
	} else {
		fmt.Println("创建失败")
	}

    if account != nil{
		account.SetAccountNo("jzh22222")
		account.SetBalance(100)
		account.SetPwd("999999")		
	}
	

	fmt.Println("account=",account,"accountNo=",account.GetAccountNo(),
	"balance=",account.GetBalance(),"pwd=",account.GetPwd())
}