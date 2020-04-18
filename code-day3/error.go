package main

//错误处理
/*
go程序内部如果有错误
通过最后一个返回值显示返回给调用者
由调用者决定如何处理
*/

import (
	"errors"
	"fmt"
	"strconv"
)

func main()  {
	value,err :=strconv.Atoi("xxx")
	fmt.Printf("%t\n",err)
	fmt.Println(err)
	fmt.Println(value)

	e:=fmt.Errorf("自定义error")
	fmt.Printf("%T  %#v\n",e,e)

	ee:=errors.New("New 自定义error")
	fmt.Printf("%T  %#v\n",ee,ee)
	statu,err:=chufa(1,0)
	fmt.Printf("%v,%v",statu,err)
}

func chufa(n1 int ,n2 int)(int,error)  {
	if n2>0 {
		return n1/n2,nil
	} else {
		return 0,errors.New("除数为0。。。")
	}
}