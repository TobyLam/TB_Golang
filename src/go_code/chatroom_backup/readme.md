介绍：海量用户通讯系统
说明：采用tcp通讯，cs结构
功能：
	1.登录
	2.注册
	3.群聊
	4.在线用户列表
运行：
	 1.进入GOPATH目录（go env 命令可查看）
     2.编译客户端 go build -o client.exe go_code/chatroom/client/main,生成客户端可执行文件client.exe
     3.编译服务端 go build -o server.exe go_code/chatroom/server/main,生成服务端可执行文件server.exe
     4.运行server.exe 启动服务端
     5.运行client.exe 启动客户端


