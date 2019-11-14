package main

import "fmt"
//一个文件夹下只能有一个包
//要改变内容必须使用指针
//结构体的定义
//为结构体定义得方法必须放在同一个包内，可以为不同的文件
type treeNode struct {
	value       int
	left, right *treeNode
}

func createTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

//给结构定义方法,传值
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

//(node *treeNode)结构体，(value int) 参数
func (node *treeNode) setValue(value int) {

	node.value = value
}

//结构创建分配不需要知道在堆上（垃圾回收）还是栈上（局部变量） 返回局部变量得地址
//只有使用指针才可以改变结构内容
//nil指针也可以调用方法
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
	root.setValue(100)
	root.print()
	root.traverse()
}
