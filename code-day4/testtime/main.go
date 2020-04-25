package main

import (
	"fmt"
	"time"
)

func main() {
	//时间函数
	now := time.Now()
	fmt.Printf("%T,%#v\n", now, now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	//获取当前uninx时间戳
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	//时间格式转换，字符串to 时间
	//Printf 格式化字符串。占位符
	//占位符
	//2006 年
	//01 月
	//02 日
	//03 12进制时
	//15 24进制小时
	//04 分
	//05 秒

	fmt.Println(now.Format("2006-01-02 03:04:05"))
	//生成时间

	//字符串转时间
	cTime, err := time.Parse("2006-01-02", "2001-03-04")
	fmt.Println(cTime, err)

	//now-ctime
	dura := time.Since(cTime)
	fmt.Println(dura)
	//ctime-now
	dura = time.Until(cTime)
	fmt.Println(dura)

	//时间函数
	dayInterval, _ := time.ParseDuration("-24h")
	fmt.Println(now, now.Add(dayInterval))
	yesterday := now.Add(dayInterval)
	fmt.Println(now, yesterday)
	fmt.Println(yesterday.After(now))
	fmt.Println(yesterday.Before(now))
	//sub计算两个时间差
	fmt.Println(yesterday.Sub(now))

	fmt.Println(time.Now())
	time.Sleep(time.Second * 3)
	fmt.Println(time.Now())
}
