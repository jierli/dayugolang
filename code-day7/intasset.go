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

type SMSSender struct {
	api string
	id  string
	key string
}

func (s SMSSender) Send(to, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil

}

func (s SMSSender) SendAll(to []string, msg string) error {
	fmt.Printf("发送邮件：%s内容： %s\n", to, msg)
	return nil

}

//输入参数为接口类型
func PrintConfig(sender Sender) {
	//需要将sender转换为结构体对象
	//断言
	obj, ok := sender.(EmailSender)
	fmt.Printf("%t,%#v", ok, obj)
}

func main() {

	emailSender := EmailSender{"smtp.qq.com", 123, "dayu", "aaaaa"}
	PrintConfig(emailSender)

	smsSender := SMSSender{"smtp.qq.com", "123123", "dayu"}
	PrintConfig(smsSender)
}
