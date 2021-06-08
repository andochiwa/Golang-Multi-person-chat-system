package main

import (
	"net"
	"redis.demo/common/message"
)

type CorrentUser struct {
	Conn net.Conn
	User message.User
}
