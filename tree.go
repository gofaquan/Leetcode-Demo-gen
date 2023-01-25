package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(nums); i += 2 {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != -1 {
			left := &TreeNode{Val: nums[i]}
			node.Left = left
			queue = append(queue, left)
		}
		if nums[i+1] != -1 {
			right := &TreeNode{Val: nums[i+1]}
			node.Right = right
			queue = append(queue, right)
		}
	}
	return root
}

func parse(s string) []int {
	var res = make([]int, 0, 16)
	s = strings.Trim(s, "[]")     // 去除字符串两端的 [ ]
	strs := strings.Split(s, ",") // 分割字符串

}

func main() {
	strs := "[7, 3, 15, null, null, 9, 20]"
	parse(strs)
	nums := []int{7, 3, 15, -1, -1, 9, 20}
	root := createBinaryTree(nums)
	fmt.Println(root)
}
