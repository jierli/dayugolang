package main

import "fmt"

func main() {
	//多维切片，和多维数组类似
	multi := [][]string{}
	fmt.Printf("%#v\n", multi)
	multi = append(multi, []string{"a", "b", "c"})
	multi = append(multi, []string{"x", "y", "z"})
	fmt.Printf("%v\n", multi)
	fmt.Printf("%v,%v\n", multi[0][0], multi[1][1])
}
