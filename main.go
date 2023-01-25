package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func buildList(s string) *ListNode {
	s = strings.Trim(s, "[]")     // 去除字符串两端的 [ ]
	strs := strings.Split(s, ",") // 分割字符串

	var dummyHead = &ListNode{}
	cur := dummyHead
	for _, str := range strs {
		n, _ := strconv.Atoi(str) // 将字符串转换为整数
		cur.Next = &ListNode{Val: n}
		cur = cur.Next
	}

	return dummyHead.Next
}

func main() {
	s := "[1,2,3,3,4,4,5]"
	head := buildList(s)
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}
}
