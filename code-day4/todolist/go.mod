module todolist

go 1.13

//go mod init "github.com/jierli/dayupaktest"
//go mod init todolist
//导入目录名
//调用包名
//包名package 与所在目录名一致

//go mod可以修改版本名
//相关命令 go list
// go mod edit -replace=
require (
	github.com/imsilence/strutil v0.0.0-20200425035852-949389850a34
	github.com/jierli/dayupaktest v0.0.0-20200425060547-0f7ec48b961a
)
