package main

import "fmt"

type EmailSender struct {
	addr     string
	prot     int
	user     string
	password string
}

func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil
}

func main() {
	//定义匿名接口
	var sender interface {
		Send(to, msg string) error
	}
	fmt.Printf("%T,%#v\n", sender, sender)
	fmt.Println(sender)
	sender = EmailSender{}
	fmt.Println(sender)
}
