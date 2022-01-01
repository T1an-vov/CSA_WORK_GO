package main

import (
	"CSA_Work_Go/QandA/proto"
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

type GetUserReq struct{
	UserID int64 `json:"userId"`
}
type GetUserResp struct {
	UserID int64 `json:"userId"`
	UserName string `json:"userName"`
}
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		msg,err:=proto.Decode(reader)
		if err ==io.EOF  {
			fmt.Println("回答结束")
			break
		}
		if err != nil &&err!=io.EOF{
			fmt.Println("read from client failed, err:", err)
			break
		}
		fmt.Println("收到来自学生的提问：", msg)
		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		b,err:=proto.Encode(inputInfo)
		if err != nil {
			fmt.Println("encode failed")
			break
		}
		_, err = conn.Write(b) // 发送数据
		if err != nil {
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}