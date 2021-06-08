package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"os"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// 登录函数
func login(id int, password string) error {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		return err
	}
	defer conn.Close()

	// 发送消息给服务器
	msg := message.Message{}
	msg.Type = message.LoginMessageType
	// 创建一个LoginMessage结构体
	loginMsg := message.LoginMessage{}
	loginMsg.UserId = id
	loginMsg.UserPwd = password

	// 将loginMsg序列化
	data, err := json.Marshal(loginMsg)
	if err != nil {
		return err
	}
	// 将序列化好的消息给msg
	msg.Data = string(data)

	// 将msg序列化，然后发送给服务器
	data, err = json.Marshal(msg)
	if err != nil {
		return err
	}

	// 发送数据
	err = utils.WritePkg(conn, data)
	if err != nil {
		return err
	}

	// 获取数据
	msg, err = utils.ReadPkg(conn)
	if err != nil {
		return err
	}
	// 将msg的data部分反序列化
	var loginResult message.LoginResult
	err = json.Unmarshal([]byte(msg.Data), &loginResult)
	if err != nil {
		return err
	}

	if loginResult.Code == 200 {
		for {
			fmt.Println("登录成功！")
			// 初始化用户在线列表
			for _, v := range loginResult.Users {
				if v.UserId == id {
					// 初始化个人信息
					currentUser = CorrentUser{Conn: conn, User: v}
					continue
				}
				InsertUser(v)
			}
			// 启动协程读取消息
			go processServerMessage(conn)
			for {
				showMenu()
			}
		}
	} else {
		fmt.Println(loginResult.Error)
	}

	return nil
}

func register(id int, password string, name string) error {
	// 连接到服务器
	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		return err
	}
	defer conn.Close()

	registerMessage := message.RegisterMessage{UserId: id, UserPwd: password, UserName: name}
	// 序列化账号密码等数据
	data, err := json.Marshal(registerMessage)
	if err != nil {
		return err
	}

	msg := message.Message{Type: message.RegisterMessageType, Data: string(data)}

	// 序列化发送的信息
	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// 发送给服务器
	err = utils.WritePkg(conn, msgData)
	if err != nil {
		return err
	}

	// 读取数据
	msg, err = utils.ReadPkg(conn)
	if err != nil {
		return err
	}
	// 反序列化
	var registerResult message.RegisterResult
	err = json.Unmarshal([]byte(msg.Data), &registerResult)
	if err != nil {
		return err
	}

	if registerResult.Code == 200 {
		fmt.Println("注册成功，可以重新登录")
	} else {
		return errors.New(registerResult.Error)
	}
	return nil
}

func showMenu() {
	fmt.Println("\t\t1. 显示在线用户列表")
	fmt.Println("\t\t2. 发送消息")
	fmt.Println("\t\t3. 消息列表")
	fmt.Println("\t\t4. 退出系统")
	fmt.Println("\t\t请选择(1-4):")
	key := 0
	for !(key >= 1 && key <= 4) {
		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			fmt.Println("输入有误，请重新输入")
		}
	}
	switch key {
	case 1:
		showOnlineUser()
	case 2:
		fmt.Println("输入你要发送的信息")
		var content string
		for _, err := fmt.Scanf("%s\n", &content); err != nil; _, err = fmt.Scanf("%s\n", &content) {
			fmt.Println("输入有误，请重新输入")
			fmt.Println("输入你要发送的信息")
		}
		err := SendMessage(content)
		if err != nil {
			fmt.Println("SendMessage err =", err)
		}
	case 3:
		err := getMessageRecordRequest()
		if err != nil {
			fmt.Println("getMessageRecord err =", err)
		}
	case 4:
		os.Exit(0)
	default:
		fmt.Println("输入有误，请重新输入")
	}
}
