package Cal

import (
	"testing"
)

func TestGetSub(t *testing.T) {
	//调用
	res := getSub(10, 5)
	if res != 5 {
		//fmt.Println("AddUpper(10) 执行错误，期望值=%v 实际值=%v\n", 55, res)
		t.Fatalf("getSub(10,5) 执行错误，期望值=%v 实际值=%v\n", 5, res)
	}

	//如果正确，输出日志
	t.Logf("getSub(10,5)执行正确...")

}
