package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	//base64包 0-9a-zA-Z+/ 64个字符
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("dayu顶顶顶顶")))
	str, _ := base64.StdEncoding.DecodeString("ZGF5demhtumhtumhtumhtg==")
	fmt.Println(string(str))

	//URl
	fmt.Println(base64.URLEncoding.EncodeToString([]byte("dayu顶顶顶顶")))
	str, _ = base64.URLEncoding.DecodeString("ZGF5demhtumhtumhtumhtg==")
	fmt.Println(string(str))
}
