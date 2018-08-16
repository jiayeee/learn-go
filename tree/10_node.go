package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 为结构体定义方法
func (node Node) Print() {
	fmt.Println(node.Value)
}

// 定义构造方法
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("setting Value to nil node, ignored !")
		return
	}
	node.Value = value
}

func (node *Node) Traverse() {
	if node == nil {
		return
	}

	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
