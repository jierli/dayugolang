package main

import "fmt"

//定义接口

type Sender interface {
	//定义接口的行为（方法名，参数列表，返回值）方法名与参数列表必须一致
	Send(string, string) error
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

func (s EmailSender) SendCc(to []string, cc string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, cc, msg)
	return nil
}

func main() {
	//定义变量
	var sender Sender
	fmt.Printf("%T %#v\n", sender, sender)

	//赋值
	//创建Emailsend对象
	emailSender := EmailSender{"stmp.qq.com", 1111, "dayu", "123"}
	sender = emailSender
	fmt.Printf("%T %#v\n", sender, sender)

	//接口对象不能访问属性
	//fmt.Println(sender.addr)
	//报错

	//接口对象调用接口方法
	sender.Send("dayu", "dayutest")
	member := []string{"dayu1", "dayu2"}
	sender.SendAll(member, "dayutest")
}
