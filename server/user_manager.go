package main

import (
	"fmt"
	"net"
)

var userManager *UserManager

type UserManager struct {
	OnlineUsers map[User]net.Conn
}

func init() {
	userManager = &UserManager{OnlineUsers: make(map[User]net.Conn)}
}

// AddOnlineUser 用户上线
func (this *UserManager) AddOnlineUser(user User, conn net.Conn) {
	this.OnlineUsers[user] = conn
}

// DeleteOnlineUser 用户下线
func (this *UserManager) DeleteOnlineUser(user User) {
	delete(this.OnlineUsers, user)
}

func (this *UserManager) GetOnlineUser(user User) (net.Conn, error) {
	conn, ok := this.OnlineUsers[user]
	if ok == false {
		return nil, fmt.Errorf("user %s not exists or is not online\n", user.UserName)
	}
	return conn, nil
}

// GetAllOnlineUser 获取所有在线用户
func (this *UserManager) GetAllOnlineUser() map[User]net.Conn {
	return this.OnlineUsers
}
