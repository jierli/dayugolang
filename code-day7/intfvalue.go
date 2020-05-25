package main

import "fmt"

type Sender interface {
	Send(string, string) error
	SendAll([]string, string) error
}

type SingleSender interface {
	Send(string, string) error
}

type EmailSender struct {
	addr     string
	prot     int
	user     string
	password string
}

//
func (s EmailSender) Send(to string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil
}

//
func (s EmailSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil
}

func main() {
	var sender Sender
	var singleSender SingleSender

	emailSender := EmailSender{}

	sender = emailSender
	singleSender = sender
	fmt.Printf("%T %#v\n", sender, sender)
	fmt.Printf("%T %#v\n", singleSender, singleSender)

	//不能赋值
	//SingleSender does not implement Sender (missing SendAll method)
	/*
		singleSender = emailSender
		sender = singleSender
		fmt.Printf("%T %#v\n", sender, sender)
		fmt.Printf("%T %#v\n", singleSender, singleSender)
	*/

}
