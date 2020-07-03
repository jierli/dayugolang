package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	)

const (
	DbDriver   = "mysql"
	DbUser     = "todolist"
	DbPassword = "todolist"
	DbName     = "todolist"
	DbHost     = "10.104.249.7"
	DbPort     = 3306
)


type Mysql struct {
	DB *sql.DB
}

func newMysql() *Mysql  {
	dsn:=fmt.Sprint("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true",DbDriver,DbUser,DbHost,DbPort,DbName)
	db,err:=sql.Open(DbDriver,dsn)
	if err != nil{
		fmt.pr
	}
}