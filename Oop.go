package main

import "fmt"

//结构体的定义
type treeNode struct {
	value       int
	left, right *treeNode
}

func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}
func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createTreeNode(2)
	fmt.Println(root.left.right)
	fmt.Println("ddddddddddd")
	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)
}
