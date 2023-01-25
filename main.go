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

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

type Parser struct {
	s    string
	data []string
}

// parseString 处理字符串多余的字符, 读取数据到 data
func (p *Parser) parseString() {
	p.s = strings.Trim(p.s, "[]")           // 去除字符串两端的 [ ]
	p.s = strings.Replace(p.s, " ", "", -1) // 去除空格
	p.data = strings.Split(p.s, ",")        // 按 , 分割得到数据切片
}

func (p *Parser) buildList() *ListNode {
	var dummyHead = &ListNode{}
	cur := dummyHead
	for _, str := range p.data {
		num, _ := strconv.Atoi(str) // 将字符串转换为整数
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}

	return dummyHead.Next
}

func main() {
	var p = Parser{}
	p.s = "[1,2,3, 3,4,4,5]"
	p.parseString()
	head := p.buildList()
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%d ", cur.Val)
	}

}
