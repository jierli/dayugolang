package main

import "os"

func CopyFile()  {
	

	srcFile,err:=os.OpenFile(src)
	if err!=nil{
		return
	}
	defer srcFile.Close()

	dstFile,err:=os.OpenFile(dst)
	if err!=nil{
		return
	}
	defer dstFile.Close()
}