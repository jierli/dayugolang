package main

import "sync"

type pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

// New 初始化连接池,设置最大并发上限
func NewPool(size int) *pool {
	if size <= 0 {
		size = 1
	}
	return &pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

// Add 添加或减少队列
func (p *pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	for i := 0; i > delta; i-- {
		<-p.queue
	}
	p.wg.Add(delta)
}

// Done 完成.释放通道
func (p *pool) Done() {
	<-p.queue
	p.wg.Done()
}
