package test

import (
	"reflect"
	"testing"
)

// 测试用例，命令 go test -v main_test.go
// 文件命名 xxx_test.go
func TestReflectFunc(t *testing.T) {
	//定义两个函数
	call1 := func(v1 int, v2 int) {
		t.Log(v1, v2)
	}

	call2 := func(v1 int, v2 int, s string) {
		t.Log(v1, v2, s)
	}

	var (
		function reflect.Value
		inValue  []reflect.Value
		n        int
	)
	//定义一个适配器函，作用：统一处理接口
	bridge := func(call interface{}, args ...interface{}) {
		n = len(args)
		inValue = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			inValue[i] = reflect.ValueOf(args[i])
		}
		function = reflect.ValueOf(call)
		function.Call(inValue)
	}

	bridge(call1, 1, 2)
	bridge(call2, 1, 2, "test2")

}
