package main

import (
	"fmt"
	"learn-go/tree"
)

/**
 * 通过组合的方式扩展系统类型或者别人的类型 Node 实现后序遍历
 *
 */
type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {

	root := tree.Node{}

	// 赋值
	root.Value = 10
	root = tree.Node{Value: 111}
	root.Left = &tree.Node{Value: 12}
	root.Right = &tree.Node{3, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(11)

	fmt.Println(root)

	fmt.Println()

	nodes := []tree.Node{
		{1, nil, nil},
		{},
		{2, nil, &root},
	}
	fmt.Println(nodes)

	root.Print()

	fmt.Println("--------------")

	root.Left.Right.SetValue(1)
	root.Left.Right.Print()

	pRoot := &root
	pRoot.Print()
	pRoot.SetValue(100)
	pRoot.Print()

	pRoot = nil
	pRoot.SetValue(200)

	fmt.Println("------------------------------")

	fmt.Println("     [3]")
	fmt.Println("    /  \\")
	fmt.Println("  [0]  [4]")
	fmt.Println("   \\    /")
	fmt.Println("   [2] [5]")

	root.SetValue(3)
	root.Left.SetValue(0)
	root.Left.Right.SetValue(2)
	root.Right.SetValue(4)
	root.Right.Left.SetValue(5)
	root.Traverse()

	fmt.Println("------------------------------")
	myNode := myTreeNode{&root}
	myNode.postOrder()
}
