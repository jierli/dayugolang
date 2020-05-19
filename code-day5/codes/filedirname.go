package main

import (
	"fmt"
	"os"
)

func dir(paht string) {
	file, err := os.Open(paht)
	if err != nil {
		return
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, name := range names {
		fpath = file + "/" + name
		if fileInfo, err := os.Stat(fpath); err == nill {
			if fileInfo.IsDir() {
				fmt.Println("dir:", fpath)
				dir(fpath)
			}
			fmt.Println("file:", fpath)
		}
	}
}

func main() {
	file, err := os.Open("todolist")
	if err != nil {
		return
	}
	defer file.Close()

	names, err := file.Readdirnames(-1)
	for i, j := range names {
		fmt.Println(i, j)
	}

}
