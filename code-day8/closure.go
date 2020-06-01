package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	//wg.Add(9)
	for i:=0;i<10;i++{
		wg.Add(1)
		go func (i int)  {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}