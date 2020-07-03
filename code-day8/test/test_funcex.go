package main

import "fmt"

type user struct {
	name string
	age  int
	addr string
}

func (user user) getuser(name string) {
	fmt.Println(user.name)
}

func main() {
	u := user{
		"dayu",
		34,
		"xiameng",
	}
	get1 := (*user).getuser
	get2 := user.getuser
	fmt.Printf("%T,%T\n", get1, get2)
	get2(u, "dayu")
	fv := u.getuser
	fmt.Printf("%T\n", fv)
}
