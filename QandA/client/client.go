package main

import (
	"CSA_Work_Go/QandA/proto"
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)
func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:8080")
	if err != nil {
		fmt.Println("client failed")
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("输入问题，按Q退出")
		input, _ := inputReader.ReadString('\n') // 读取用户输入
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		b,err:=proto.Encode(inputInfo)
		if err != nil {
			fmt.Println("encode failed")
		}
		_, err = conn.Write(b) // 发送数据
		if err != nil {
			return
		}
		reader:=bufio.NewReader(conn)
		msg,err:=proto.Decode(reader)//将收到的数据解码
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("收到老师的回答："+msg)
	}
}