package main

import "fmt"

var users [10]string=[10]string{"dayu","xiamen","dba","dayu1","xiamen1","dba2","test"}

func main()  {
	bounds:=[...]int{1,2,3,4,5,6,7}
	fmt.Println(users,bounds)
	fmt.Println(bounds==[...]int{1,2,3,4,5,6,7})
	fmt.Println(bounds[2:5])
	fmt.Printf("%T %q\n",bounds[2:5],bounds[2:5])
	fmt.Printf("%T %q\n",users[2:5],users[2:5])
	fmt.Printf("%T %q\n",users[2:5:6],users[2:5:6])
}