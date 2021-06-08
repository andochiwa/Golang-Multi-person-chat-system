package main

import (
	"fmt"
	"redis.demo/common/message"
)

var currentUser CorrentUser
var OnlineUsers = make(map[int]message.User)

// InsertUser 添加在线用户
func InsertUser(user message.User) {
	OnlineUsers[user.UserId] = user
}

// deleteUser 删除离线用户
func deleteUser(userId int) {
	delete(OnlineUsers, userId)
}

// showOnlineUser 显示当前用户
func showOnlineUser() {
	fmt.Println("Online user list:")
	for _, v := range OnlineUsers {
		fmt.Printf("id: %d, name: %s\n", v.UserId, v.UserName)
	}
}
