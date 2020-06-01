package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func fileLine(path string) int  {
	cnt:=0
	file,err:=os.Open(path)
	if err!=nil{
		return cnt
	}
	defer file.Close()
	reader:=bufio.NewReader(file)
	for {
		ctx,_,err:=reader.ReadLine()
		if err!=nil{
			break
		}
		if strings.TrimSpace(string(ctx))==""{
			continue
		}
		cnt++
	}
	return cnt
	
}

func main()  {
	fmt.Println(fileLine("fileLine.go"))
}