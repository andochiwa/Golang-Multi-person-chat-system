package main

import (
	"encoding/json"
	"fmt"
	"net"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

// 和服务器保持通讯
func processServerMessage(conn net.Conn) {
	for {
		msg, err := utils.ReadPkg(conn)
		if err != nil {
			fmt.Println("读取消息失败 =", err)
			return
		}
		switch msg.Type {
		case message.NotifyUserStatusType:
			// 有人上线或离线了
			var user message.User
			err := json.Unmarshal([]byte(msg.Data), &user)
			if err != nil {
				fmt.Println("processServerMessage json.Unmarshal err =", err)
			} else {
				switch user.Status {
				case message.UserOnline:
					InsertUser(user)
				case message.UserOffline:
					deleteUser(user.UserId)
				default:
					fmt.Println("can not resolve user")
				}
			}
		case message.SmsResultType:
			// 收到消息
			var smsMessage message.SmsMessage
			err := json.Unmarshal([]byte(msg.Data), &smsMessage)
			if err != nil {
				fmt.Println("processServerMessage json.Unmarshal err =", err)
				break
			}
			fmt.Printf("收到用户%s消息: %s\n", smsMessage.UserName, smsMessage.Content)
		case message.SmsRecodeResultType:
			// 获取消息记录
			err := getMessageRecordResponse(&msg)
			if err != nil {
				fmt.Println("getMessageRecordResponse err =", err)
				break
			}
		default:
			fmt.Println("收到消息 =", msg.Data)
		}
	}
}
