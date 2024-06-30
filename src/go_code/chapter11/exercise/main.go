package main

type Ainterface interface {
	Test01()
	Test02()
}

type Binterface interface {
	Test01()
	Test03()
}

type Cinterface interface {
	Ainterface
	Binterface //低版本不支持同名方法，会报错
}

func main() {

}
