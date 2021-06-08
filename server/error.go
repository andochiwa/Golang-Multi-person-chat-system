package main

import "errors"

var (
	ErrorUserNotExists = errors.New("user not exists")
	ErrorUserExists    = errors.New("user already exists")
	ErrorUserPwd       = errors.New("incorrect password")
)
