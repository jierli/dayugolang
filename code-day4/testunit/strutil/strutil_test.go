package strutil

import "testing"

//基准测试
//go test -run none -bench=.
//go test -run none -bench=. -benchmem ./... 增加内存分配分析
/*
goos: windows
goarch: amd64
pkg: testunit/strutil
BenchmarkString2Intv1-8         1000000000               0.00598 ns/op
BenchmarkString2Intv2-8         1000000000               0.0130 ns/op
PASS
ok      testunit/strutil        0.411s
*/
func BenchmarkString2Intv1(b *testing.B) {
	for i := 0; i < 100000; i++ {
		String2Intv1(i)
	}
}

func BenchmarkString2Intv2(b *testing.B) {
	for i := 0; i < 100000; i++ {
		String2Intv2(i)
	}
}
