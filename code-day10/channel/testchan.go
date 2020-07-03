package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	channel := make(chan int, 10)
	go func(){
		for i := 0; i < 10; i++ {
			a :=<- channel
			fmt.Println( a,"<-")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("<-", i)
			channel <- i
		}
		//close(channel)
		wg.Done()
	}()
	wg.Add(2)
	//func() {
	//	for e := range channel {
	//		fmt.Println(e)
	//	}
		//wg.Done()
	//}()
	wg.Wait()


}
