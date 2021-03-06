package main

import "fmt"

type Sender interface {
	Send(string, string) error
	Send1(string, string) error
	Send2(string, string) error
	SendAll([]string, string) error
	Send11(string) error
}

type SingleSender interface {
	Send(string, string) error
}

type SingleSender1 interface {
	Send11(string) error
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
func (s EmailSender) Send11(to string) error {
	fmt.Printf("发送邮件：%s\n", to)
	return nil
}
func (s EmailSender) Send1(to string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil
}
func (s EmailSender) Send2(to string, msg string) error {
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
	//singleSender = sender
	//fmt.Printf("%T %#v \n", sender, sender)
	//fmt.Printf("%T %#v\n", singleSender, singleSender)

	emailSender := EmailSender{}
	fmt.Printf("%T %#v \n", emailSender, emailSender)

	sender = emailSender
	//sender1 = emailSender
	//sender2 = emailSender
	singleSender = sender
	fmt.Printf("%T %#v \n", sender, sender)
	fmt.Printf("%T %#v\n", singleSender, singleSender)
	fmt.Println(sender.Send1("dayu1", "dayu2"))
	fmt.Println(sender.Send2("dayu1", "dayu2"))

	//不能赋值
	//SingleSender does not implement Sender (missing SendAll method)
	/*
		singleSender = emailSender
		sender = singleSender
		fmt.Printf("%T %#v\n", sender, sender)
		fmt.Printf("%T %#v\n", singleSender, singleSender)
	*/

	var singleSender1 SingleSender1
	singleSender1 = sender
	fmt.Printf("%T %#v\n", singleSender1, singleSender1)
	singleSender1.Send11("dayutest")

	value, ok := singleSender1.(EmailSender)
	fmt.Printf("%T,    %#v,    %v", value, value.prot, ok)

}
