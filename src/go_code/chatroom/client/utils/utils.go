package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

//细节说明：
/*
1.网络传输可能出现丢包问题，故约定 先发送数据长度（字节数n）,再发送数据本身
1.1 数据本身可以携带数据长度，即一次性发送数据长度和数据本身，发送前计算出发送数据长度，再与数据一起发送，由于本项目使用的是tcp长连接，可以先发送长度，再发送数据本身。而类似网站的http协议，一般为短连接，一次性读写完成后立即关闭，就需要一次性发送数据长度及数据本身
2.网络传输长度不能使用int类型，需要使用字节byte，所以需要进行转换
3.连接不会主动中断，故接收方先接收数据长度，再根据该长度再次接收数据本身
*/

// 将方法关联到结构体中
type Transfer struct {
	//字段
	Conn net.Conn //连接
	//定义一个字节数组，数组大小根据实际情况判断，8096个字节的数据
	Buf [8096]byte //这是传输时，使用缓冲
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("正在读取服务器返回的数据...")

	//Conn.Read(this.Buf[:n]) 解析：从Conn这个连接里面读取n个字节，赋值给this.Buf;并非从this.Buf里读取数据。

	//1.读取数据长度
	//conn.Read 在conn没有被关闭的情况下，才会阻塞,如果客户端（发送数据方）关闭了 conn 则不会阻塞
	//约定发送的数据大小是4个字节【unit32类型】，故读取也是4个字节

	//未读前的长度：
	//fmt.Println("读取数据长度前，当前的buf=", this.Buf[:4])
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	//读到的长度：
	//fmt.Println("读取数据长度时，读到的buf=", this.Buf[:4])

	//2.数据类型转换
	//读到的长度类型为byte切片，需要进行转换
	//根据buf[:4] 转换成一个 uint32类型，直观的数据实际类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//fmt.Println("转换后的数据长度，pkgLen=", pkgLen)

	//3.根据pkgLen读取消息内容
	//fmt.Println("读取消息前，当前的buf=", this.Buf[:pkgLen])
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	//校验 n 和 期望读取的数据长度pkgLen 是否相等
	if n != int(pkgLen) || err != nil {
		return
	}
	//fmt.Println("读取数据后，读到的buf=", this.Buf[:pkgLen])

	//把this.Buf[:pkgLen] 反序列化成  -> message.Message
	//mes需要加&，否则返回的mes为空！！
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//fmt.Println("反序列化后的数据，mes=", mes)
	return

}

func (this *Transfer) WritePkg(data []byte) (err error) {

	//1. 先发送一个长度（pkgLen）给对方

	//unit32无符号，可发送2的32次方个字节的数据，大约4G 的内容【2的32次方/1024/1024/1024】
	var pkgLen uint32

	//conn.Write（）只支持字节传输，故需要转换
	//len（）先获取data的长度，类型为int
	//uint32()进行强制转换
	pkgLen = uint32(len(data))
	//unit32类型 32位只需要4个字节存储，因此声明大小为4的byte切片足够
	//var buf [4]byte ，直接使用Transfer的Buf，取4个字节即可
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)

	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//2. 发送data本身,要求data是[]byte类型
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
