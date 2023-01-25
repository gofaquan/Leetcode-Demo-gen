package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "[7, 3, 15, null, null, 9, 20]"
	fmt.Println(s)
	s = strings.Trim(s, "[]") // 去除字符串两端的 [ ]
	//strs := strings.Split(s, ",") // 分割字符串
	trimSpaceS := strings.Replace(s, " ", "", -1)

	for _, i2 := range strings.Split(trimSpaceS, ",") {
		//fmt.Println(i)
		fmt.Println(i2)
	}

}
