package main

import "fmt"

type BinaryNode struct {
	data int64
	left *BinaryNode
	right *BinaryNode
}

type BST struct {
	root *BinaryNode
}

func (t *BST) insert(data int64) {
	if t.root == nil {
		t.root = &BinaryNode{data:data, left: nil, right: nil}
		return
	}
	t.root.insert(data)
}

func (node *BinaryNode) insert(data int64) {
	if node == nil {
		return
	}

	if node.data >= data {
		if node.left == nil {
			node.left = &BinaryNode{data:data, left:nil, right:nil}
		} else {
			node.left.insert(data)
		}
	} else {
		if node.right == nil {
			node.right = &BinaryNode{data:data, left:nil, right:nil}
		} else {
			node.right.insert(data)
		}
	}
}

func (t BST) leftOrderTraversal() {
	if t.root == nil {
		return
	}
	fmt.Printf("Left order trarversal : ")
	t.root.leftOrderTraversal()
	fmt.Printf("\n")
}


func (node *BinaryNode) leftOrderTraversal() {
	if node == nil {
		return
	}

	node.left.leftOrderTraversal()
	fmt.Printf("%d ", node.data)
	node.right.leftOrderTraversal()

}


func main() {
	fmt.Printf("Binary search tree\n")
	bst := &BST{nil}

	bst.insert(10)
	bst.insert(1)
	bst.insert(12)
	bst.insert(15)
	bst.insert(2)
	bst.insert(5)

	bst.leftOrderTraversal()
}
