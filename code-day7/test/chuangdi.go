package main

import "fmt"

 


func main()  {
	e1,e2:="dayu",[]string{"dayu","dba"}
	fmt.Printf("e1:%p %v\n",&e1,e1)
	func (e string)  {
		fmt.Printf("e:%p %v\n",&e,e)
		e="silence"
	}(e1)
 

	fmt.Printf("e2:%p %v\n",&e2,e2)
	func (e []string)  {
		fmt.Printf("e:%p %v\n",&e,e)
		e[1]="woniub"
	}(e2)
	fmt.Printf("e2:%p %p %s %v\n",&e1,&e2,e1,e2)
 

}