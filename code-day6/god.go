package main

import (
	"encoding/gob"
	"os"
	"time"
)

type Task struct {
	id        int
	name      string
	status    int
	startTime *time.Time
	endTime   *time.Time
	user      string
}

func main() {

	task := &Task{{
		id:        1,
		name:      "整理课程笔记",
		startTime: &now,
		endTime:   &end,
		user:      "dayu",
	}, {
		id:        2,
		name:      "整理课程笔记2",
		startTime: &now,
		endTime:   &end,
		user:      "dayu",
	}}

	file, _ := os.Create("./test/user.gob")
	defer file.Close()

	encoder := gob.Encoder(task)
}
