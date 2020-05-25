package main

import "fmt"

type SingleSender interface {
	Send(string, string) error
}

type Sender interface {
	SingleSender
	SendAll([]string, string) error
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

func (s EmailSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil
}

func main() {
	var sender Sender = EmailSender{}
	fmt.Println(sender.Send)
}
