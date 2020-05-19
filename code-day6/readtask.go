package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sqs/goreturns/returns"
)

type Task struct {
	id        int
	name      string
	status    int
	startTime *time.Time
	endTime   *time.Time
	user      string
}
func ParseTask(line string) (*Task ,error) {
	node,_:=string.Split(line,",")
	if len(node)!==6{
		return errors.New("数据不正确")
	}
	id,err:=strconv.Atoi(nodes[0])
	if err!=nil{
	return nil，err
}
}

func main()  {
	file,_:=os.Open("./test/user.txt")
	defer file.Close()
	task:=make([]*Task,0, 100)
	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		task:=ParseTask(scanner.Text())
	}
}