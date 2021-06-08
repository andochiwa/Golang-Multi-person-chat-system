package main

import (
	"fmt"
	"os"
)

func main() {
	// 接收用户的选择
	key := 0
	for {
		fmt.Println("--------------欢迎登陆多人聊天系统----------------")
		fmt.Println("\t\t 1 登录聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择 (1-3):")

		_, err := fmt.Scanf("%d\n", &key)
		if err != nil {
			fmt.Println("输入有误 请重新输入")
			continue
		}
		var id int
		var password string
		var name string
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			for {
				fmt.Println("输入用户id")
				for _, err := fmt.Scanf("%d\n", &id); err != nil; _, err = fmt.Scanf("%d\n", &id) {
					fmt.Println("输入格式有误 请重新输入")
					fmt.Println("输入用户id")
				}

				fmt.Println("输入用户密码")
				for _, err := fmt.Scanf("%s\n", &password); err != nil; _, err = fmt.Scanf("%s\n", &password) {
					fmt.Println("输入格式有误 请重新输入")
					fmt.Println("输入用户密码")
				}

				err = login(id, password)
				if err != nil {
					fmt.Println("登录发生了错误 :", err)
					continue
				}
				break
			}
		case 2:
			fmt.Println("注册用户")
			fmt.Println("输入注册id")
			for _, err := fmt.Scanf("%d\n", &id); err != nil; _, err = fmt.Scanf("%d\n", &id) {
				fmt.Println("输入格式有误 请重新输入")
				fmt.Println("输入注册id")
			}

			fmt.Println("输入注册密码")
			for _, err := fmt.Scanf("%s\n", &password); err != nil; _, err = fmt.Scanf("%s\n", &password) {
				fmt.Println("输入格式有误 请重新输入")
				fmt.Println("输入注册密码")
			}

			fmt.Println("输入用户名")
			for _, err := fmt.Scanf("%s\n", &name); err != nil; _, err = fmt.Scanf("%s\n", &name) {
				fmt.Println("输入格式有误 请重新输入")
				fmt.Println("输入用户名")
			}

			err := register(id, password, name)
			if err != nil {
				fmt.Println("注册发生错误:", err)
				continue
			}
		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入不正确，请重新输入")
			continue
		}
	}
}
