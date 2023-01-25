# Leetcode-Demo-gen

编写 leetcode 代码生成工具来辅助使用 go 来刷题

Todo：

- [ ] 无错误校验机制，传入的 string 必须是正确的

考虑 leetcode debug 官方会给出 string，就还没增加错误校验的机制


### 使用示例
```go
package main

import (
	"fmt"
	"github.com/gofaquan/Leetcode-Demo-gen"
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

func main() {
	head := &ListNode{}
	parser.AtoList("[1, 2, 3, 3, 4, 5]", head)
	fmt.Println(head)

	root := &TreeNode{}
	parser.AtoBinaryTree("[7, 3, 15, null, null, 9, 20]", root)
	fmt.Println(root)
}
```
