package main

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"redis.demo/common/message"
)

type MessageDao struct {
	RedisPool *redis.Pool
}

var MyMessageDao *MessageDao

func initMessageDao() {
	MyMessageDao = &MessageDao{RedisPool: RedisPool}
}

// 保存消息记录
func (this *MessageDao) saveMessage(smsMessage message.SmsMessage) error {
	conn := this.RedisPool.Get()
	defer conn.Close()
	smsData, err := json.Marshal(smsMessage)
	if err != nil {
		return err
	}
	_, err = conn.Do("LPush", "char_record", string(smsData))
	if err != nil {
		return err
	}
	return nil
}

// 获取10条消息记录
func (this *MessageDao) getMessages() ([]message.SmsMessage, error) {
	conn := this.RedisPool.Get()
	defer conn.Close()
	result, err := redis.Values(conn.Do("lrange", "char_record", 0, 10))
	if err != nil {
		return nil, err
	}
	var smsMessages []message.SmsMessage
	for _, v := range result {
		var tempSmsMessages message.SmsMessage
		err = json.Unmarshal(v.([]byte), &tempSmsMessages)
		smsMessages = append(smsMessages, tempSmsMessages)
	}

	//fmt.Printf("smsData = %s, smsData type = %T\n", smsData, smsData)

	if err != nil {
		return nil, err
	}
	return smsMessages, nil
}
