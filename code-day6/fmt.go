package main

import (
	"fmt"
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

func timeToSting(t *time.Time) string {
	return t.Format("2006-01-01 15:04:06")
}

func main() {
	now := time.Now()
	//end := time.Now.Add(time.Hour * 24)
	end := time.Now()
	task := &Task{
		id:        1,
		name:      "整理课程笔记",
		startTime: &now,
		endTime:   &end,
		user:      "dayu",
	}
	tasks := []*Task{task}
	file, _ := os.Create("./test/user.txt")
	defer file.Close()

	for _, task := range tasks {
		file.WriteString(
			fmt.Sprintf("%d,%s,%d,%s,%s,%s\n",
				task.id,
				task.name,
				task.status,
				timeToSting(task.startTime),
				timeToSting(task.endTime),
				task.user,
			),
		)
	}

}
