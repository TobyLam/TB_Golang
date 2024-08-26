package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"net"
)

// 将方法关联到结构体中
type Transfer struct {
	//分析有哪些字段
	Conn net.Conn   //连接
	Buf  [8096]byte //这是传输时，使用缓冲

}

/**
 * 接收数据
 */
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("读取客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}

	//fmt.Println("读到的buf=", buf[:4])

	//根据buf[:4] 转换成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	//把this.Buf[:pkgLen] 反序列化成  -> message.Message
	//mes需要加&，否则返回的mes为空！！
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return

}

/**
 * 发送数据
 */
func (this *Transfer) WritePkg(data []byte) (err error) {

	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))

	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
