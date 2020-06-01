package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var wg sync.WaitGroup
	var locker sync.Mutex
	var a, b = 10000, 10000
	var count = 100
	wg.Add(2)
	//wg.Add(9)
	go func() {

		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if b > money {
				locker.Lock()
				b -= money
				time.Sleep(time.Microsecond)
				a += money
				locker.Unlock()
			}
		}
		wg.Done()
	}()
	go func() {

		for i := 0; i < count; i++ {
			money := rand.Intn(100)
			if b > money {
				locker.Lock()
				b -= money
				time.Sleep(time.Microsecond)
				a += money
				locker.Unlock()
			}
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Printf("%v,%v,%v", a, b, a+b)
}
