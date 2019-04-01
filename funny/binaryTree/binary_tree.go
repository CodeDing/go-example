package main

import "fmt"

type TreeNode struct {
	val         interface{}
	left, right *TreeNode
}

func InitTree() *TreeNode {
	return &TreeNode{val: nil, left: nil, right: nil}
}

func CreateTreeNode(val interface{}) *TreeNode {
	return &TreeNode{val: val, left: nil, right: nil}
}

func FindTreeHeight(t *TreeNode) int {
	if t == nil {
		return 0
	}
	//FindTreeHeight(l.left) > FindTreeHeight(l.right) ? FindTreeHeight(l.left) + 1 : FindTreeHeight(l.right) + 1
	var height, h_left, h_right int
	if t.left != nil {
		h_left = FindTreeHeight(t.left)
	}
	if t.right != nil {
		h_right = FindTreeHeight(t.right)
	}

	if h_left < h_right {
		height = h_right + 1
	} else {
		height = h_left + 1
	}
	return height
}

func PreOrderTraverseTree(t *TreeNode, printVal *string) {
	if t != nil {
		*printVal = fmt.Sprintf("%s->%v", *printVal, t.val)
		PreOrderTraverseTree(t.left, printVal)
		PreOrderTraverseTree(t.right, printVal)
	}
}

func InOrderTraverseTree(t *TreeNode, printVal *string) {
	if t != nil {
		InOrderTraverseTree(t.left, printVal)
		*printVal = fmt.Sprintf("%s->%v", *printVal, t.val)
		InOrderTraverseTree(t.right, printVal)
	}
}

func PostOrderTraverseTree(t *TreeNode, printVal *string) {
	if t != nil {
		PostOrderTraverseTree(t.left, printVal)
		PostOrderTraverseTree(t.right, printVal)
		*printVal = fmt.Sprintf("%s->%v", *printVal, t.val)
	}
}

func FindMaxDistance(t *TreeNode, max *int) int {
	if t == nil {
		return -1
	}
	nHeightOfLeftTree := FindMaxDistance(t.left, max) + 1
	nHeightOfRightTree := FindMaxDistance(t.right, max) + 1
	nDistance := nHeightOfLeftTree + nHeightOfRightTree
	if *max < nDistance {
		*max = nDistance
	}
	if nHeightOfLeftTree > nHeightOfRightTree {
		return nHeightOfLeftTree
	} else {
		return nHeightOfRightTree
	}

}

/*
          1
		 / \
		2   3
         \  \
		 4   6
		 /   /
		5   7
		    \
		    8
*/

func main() {

	tree := InitTree()
	node1 := CreateTreeNode(1)
	node2 := CreateTreeNode(2)
	node3 := CreateTreeNode(3)
	node4 := CreateTreeNode(4)
	node5 := CreateTreeNode(5)
	node6 := CreateTreeNode(6)
	node7 := CreateTreeNode(7)
	node8 := CreateTreeNode(8)
	node1.left = node2
	node1.right = node3
	node2.right = node4
	node4.left = node5
	node3.right = node6
	node6.left = node7
	node7.right = node8
	tree = node1

	height := FindTreeHeight(tree)
	fmt.Printf("Tree height: %d\n", height)

	s := `
  1
 / \
2   3
 \  \
 4   6
 /   /
5   7
    \
    8
	`
	fmt.Printf("Tree: %s\n", s)
	fmt.Println("Pre order tree: ")
	var preOrder string
	PreOrderTraverseTree(tree, &preOrder)
	fmt.Printf("%s\n", preOrder[2:])
	var inOrder string
	fmt.Println("In order tree: ")
	InOrderTraverseTree(tree, &inOrder)
	fmt.Printf("%s\n", inOrder[2:])
	fmt.Println("Post order tree: ")
	var postOrder string
	PostOrderTraverseTree(tree, &postOrder)
	fmt.Printf("%s\n", postOrder[2:])

	var max int
	FindMaxDistance(tree, &max)
	fmt.Printf("Max distance: %d\n", max)
}
