package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	//hash算法 --》签名时用
	fmt.Printf("%x\n", md5.Sum([]byte("吞吞吐吐")))

	//如果数据太大需要分批计算
	hasher := md5.New()
	hasher.Write([]byte("吞吞"))
	hasher.Write([]byte("吐吐"))
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

	//加盐哈希算法+md5
}
