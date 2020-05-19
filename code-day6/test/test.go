package main

import "fmt"

func FormatSize(size int64) string {
	//取余
	fsize := float64(size)
	unit := float64(1024)
	Msize := map[int]string{0: "B", 1: "KB", 2: "MB", 3: "GB", 4: "TB", 5: "TB"}
	index := 0
	for fsize > unit {
		fsize /= unit
		index++
	}
	return fmt.Sprintf("%.2f%s", fsize, Msize[index])
}

func main() {
	fmt.Println(FormatSize(1024))
	fmt.Println(FormatSize(1024 * 1024))
}
