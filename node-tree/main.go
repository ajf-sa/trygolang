package main

import (
	"errors"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Insert(value int) error {
	if t == nil {
		return errors.New("Tree is nil")
	}

	if t.Val == value {
		return errors.New("This node value already exists")
	}

	if t.Val > value {
		if t.Left == nil {
			t.Left = &TreeNode{Val: value}
			return nil
		}
		return t.Left.Insert(value)
	}

	if t.Val < value {
		if t.Right == nil {
			t.Right = &TreeNode{Val: value}
			return nil
		}
		return t.Right.Insert(value)
	}
	return nil
}

func (t *TreeNode) FindMin() int {
	if t.Left == nil {
		return t.Val
	}
	return t.Left.FindMin()
}

func (t *TreeNode) FindMax() int {
	if t.Right == nil {
		return t.Val
	}
	return t.Right.FindMax()
}

func (t *TreeNode) PrintInorder() {
	if t == nil {
		return
	}

	t.Left.PrintInorder()
	fmt.Print(t.Val)
	t.Right.PrintInorder()
}
func dfs(root *TreeNode, value int) {
	if root == nil {
		return
	}
	if root.Val == value {
		fmt.Println("Found")
		return
	}
	dfs(root.Left, value)
	dfs(root.Right, value)
}
func bfs(root *TreeNode, value int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == value {
			fmt.Println("Found")
			return
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
}

func main() {
	nt := TreeNode{}
	nt.Insert(9999)
	nt.Insert(4)
	nt.Insert(1)
	nt.Insert(2)
	nt.Insert(3)
	fmt.Println(nt.FindMin())
	fmt.Println(nt.FindMax())
	nt.PrintInorder()

}
