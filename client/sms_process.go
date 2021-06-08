package main

import (
	"encoding/json"
	"fmt"
	"redis.demo/common/message"
	"redis.demo/common/utils"
)

func SendMessage(content string) error {
	var msg message.Message
	msg.Type = message.SmsMessageType

	var smsMessage message.SmsMessage
	smsMessage.User = currentUser.User
	smsMessage.Content = content
	smsData, err := json.Marshal(smsMessage)
	if err != nil {
		return err
	}

	msg.Data = string(smsData)
	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = utils.WritePkg(currentUser.Conn, msgData)
	if err != nil {
		return err
	}
	return nil
}

func getMessageRecordRequest() error {
	msg := message.Message{Type: message.SmsRecordType}
	msgData, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = utils.WritePkg(currentUser.Conn, msgData)
	if err != nil {
		return err
	}
	return nil
}

func getMessageRecordResponse(msg *message.Message) error {
	var records message.SmsRecord
	err := json.Unmarshal([]byte(msg.Data), &records)
	if err != nil {
		return err
	}
	for _, v := range records.Records {
		fmt.Printf("用户%s: %s\n", v.UserName, v.Content)
	}
	return nil
}
