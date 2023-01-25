package parser

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Parser interface {
	parseString()
	buildList() *ListNode
	buildBinaryTree() *TreeNode
}

var _ Parser = &parser{}

type parser struct {
	s    string
	data []string
}

// parseString 处理字符串多余的字符, 读取数据到 data
func (p *parser) parseString() {
	p.s = strings.Trim(p.s, "[]")           // 去除字符串两端的 [ ]
	p.s = strings.Replace(p.s, " ", "", -1) // 去除空格
	p.data = strings.Split(p.s, ",")        // 按 , 分割得到数据切片
}

// AtoTreeNode 把对应下标的元素转换成 TreeNode
func (p *parser) AtoTreeNode(index int) *TreeNode {
	if p.data[index] == "null" {
		return nil
	}

	num, _ := strconv.Atoi(p.data[index]) // 将字符串转换为整数
	return &TreeNode{Val: num}
}

// AtoListNode 把对应下标的元素转换成 ListNode
func (p *parser) AtoListNode(index int) *ListNode {
	num, _ := strconv.Atoi(p.data[index]) // 将字符串转换为整数
	return &ListNode{Val: num}
}

// buildList 把 data 转换成 链表
func (p *parser) buildList() *ListNode {
	var dummyHead = &ListNode{}
	cur := dummyHead
	for i := range p.data {
		cur.Next = p.AtoListNode(i)
		cur = cur.Next
	}

	return dummyHead.Next
}

// buildBinaryTree 把 data 转换成 二叉树
// 就是把层序遍历的结果转换成二叉树
func (p *parser) buildBinaryTree() *TreeNode {
	// 代表 "" 空字符串
	if len(p.data) == 1 || p.data[0] == "" {
		return nil
	}
	root := p.AtoTreeNode(0)

	queue := []*TreeNode{root}
	for i := 1; i < len(p.data); i += 2 {
		node := queue[0]
		queue = queue[1:]
		left, right := p.AtoTreeNode(i), p.AtoTreeNode(i+1)
		if left != nil {
			queue = append(queue, left) // 非空就有下一层, 把节点加入队列
		}

		if right != nil {
			queue = append(queue, right)
		}

		node.Left = left // 赋值, 如果 left 本来是 nil 赋值也不会影响
		node.Right = right
	}
	return root
}

func initParser(s string) *parser {
	p := &parser{s: s}
	p.parseString()
	return p
}

// AtoList 返回链表
func AtoList(s string, target any) {
	p := initParser(s)
	convert(p.buildList(), target)
}

// AtoBinaryTree 返回二叉树
func AtoBinaryTree(s string, target any) {
	p := initParser(s)
	convert(p.buildBinaryTree(), target)
}

// convert 把 当前的代码值赋值给参数
// 因为 Leetcode 的 ListNode 和 TreeNode 结构是一样的
// 这样把值赋给用户创建的 Node 就好写题解函数了
// eg. func sortList(head *ListNode) *ListNode {}
// 否则题解函数入参返回参就是 func sortList(head *parser.ListNode) *parser.ListNode {}
// 就很不方便
func convert(src, dst any) {
	data, err := json.Marshal(src)
	if err != nil {
		fmt.Println("json Marshal failed, err: ", err)
	}

	err = json.Unmarshal(data, dst)
	if err != nil {
		fmt.Println("json Unmarshal failed, err: ", err)
	}
}
