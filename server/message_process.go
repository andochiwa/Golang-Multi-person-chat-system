package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// 获取并处理消息
func process(conn net.Conn) {
	defer conn.Close()
	var user User
	// 读取客户端发送的信息
	for {
		mes, err := utils.ReadPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("对方关闭了连接，服务正常退出")
				userManager.DeleteOnlineUser(user)
				fmt.Println(user)
				NotifyUsers(user.UserId, user.UserName, message.UserOffline)
				return
			}
			userManager.DeleteOnlineUser(user)
			fmt.Println(user)
			NotifyUsers(user.UserId, user.UserName, message.UserOffline)
			fmt.Println("readPkg err =", err)
			return
		}
		if user == (User{}) {
			user, err = serverProssMessage(conn, &mes)
		} else {
			_, err = serverProssMessage(conn, &mes)
		}
		if err != nil {
			fmt.Println("serverProcessMessage err =", err)
		}
	}
}

// 根据消息种类分发函数
func serverProssMessage(conn net.Conn, mes *message.Message) (user User, err error) {
	switch mes.Type {
	case message.LoginMessageType:
		// 处理登录
		user, err = ServerProcessLogin(conn, mes)
		if err != nil {
			fmt.Println("ServerProcessLogin err")
			return
		}
	case message.RegisterMessageType:
		// 处理注册
		err = ServerProcessRegister(conn, mes)
		if err != nil {
			fmt.Println("ServerProcessRegister err")
			return
		}
	case message.SmsMessageType:
		// 处理用户发送消息
		err = SendMessageToUsers(mes)
		if err != nil {
			fmt.Println("SendMessageToUsers err")
			return
		}
	case message.SmsRecordType:
		// 给用户返回消息记录
		smsMessages, err := MyMessageDao.getMessages()
		if err != nil {
			fmt.Println("MyMessageDao.getMessages err")
			return User{}, err
		}
		records := message.SmsRecord{Records: smsMessages}
		recordsData, err := json.Marshal(records)
		if err != nil {
			fmt.Println("json.Marshal recordsData err")
			return User{}, err
		}

		msg := message.Message{Type: message.SmsRecodeResultType, Data: string(recordsData)}
		msgData, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("json.Marshal msgData err")
			return User{}, err
		}
		err = utils.WritePkg(conn, msgData)
		if err != nil {
			return User{}, err
		}
	default:
		err = errors.New("消息类型不存在，无法处理")
		return
	}
	return
}
