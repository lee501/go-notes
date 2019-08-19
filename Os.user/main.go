package main

import (
	"fmt"
	"os/user"
)

/*
	os/user包提供User的结构体
	type User struct {
		//用户id
		Uid string
		//组id
		Gid string
		//用户名
		Username string
		//组名
		Name string
		// 用户名路径
		HomeDir string
	}

	type Group struct {
		Gid  string // group ID
		Name string
	}

func Current() (*User, error)
func Lookup(username string) (*User, error)
func LookupId(uid string) (*User, error)
*/
func main()  {
	u, _ := user.Current()
	fmt.Println(u.Uid)
	fmt.Println(u.Username)
	fmt.Println(u.HomeDir)
	fmt.Println(u.GroupIds())
}