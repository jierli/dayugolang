package main

import "fmt"

type Sender interface {
	Send(string, string) error
}

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

type SMSSender struct {
	addr     string
	prot     int
	user     string
	password string
}

func (s SMSSender) Send(to string, msg string) error {
	fmt.Printf("发送短信：%s内容： %s\n", to, msg)
	return nil
}

func main() {
	var sender Sender
	sender = SMSSender{}
	sender.Send("dayu", "sms") //调用SMSSender

	sender = EmailSender{}
	sender.Send("dayu", "email") //调用EmailSender

	//实现了不同类型数据 赋值给同一个变量
}
