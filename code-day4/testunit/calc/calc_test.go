package calc

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if 3 != Add(1, 2) {
		t.Error("1+2!=3")
	}
	//a := Add(1, 2)
	//t.Log(a)
}

//单元测试
//查看单元测试覆盖率
//go test -coverprofile=cover.out ./...
//输出html  go tool cover -html cover.out
