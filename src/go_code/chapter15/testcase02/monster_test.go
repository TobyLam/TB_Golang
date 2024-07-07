package monster

import (
	"testing"
)

// 测试用例，测试 Store方法
func TestStore(t *testing.T) {

	//先创建一个Monster实例
	monster := &Monster{
		Name:  "孙悟空",
		Age:   500,
		Skill: "七十二变",
	}

	res := monster.Store()

	if !res {
		t.Fatalf("monster.Store() 错误，希望=%v 实际为=%v", true, res)
	}

	t.Logf("monster.Store() 测试成功")
}

func TestReStore(t *testing.T) {

	//先创建一个Monster实例,不需要指定字段的值
	var monster Monster

	res := monster.ReStore()

	if !res {
		t.Fatalf("monster.ReStore() 错误，希望=%v 实际为=%v", true, res)
	}

	//进一步判断
	if monster.Name != "孙悟空" {
		t.Fatalf("monster.ReStore() 错误，希望=%v 实际为=%v", "孙悟空", monster.Name)
	}

	t.Logf("monster.ReStore() 测试成功")
}
