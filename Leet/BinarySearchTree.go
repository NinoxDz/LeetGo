package Leet

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type BinaryNode struct {
	left  *BinaryNode
	right *BinaryNode
	data  int64
}

type BinaryTree struct {
	root *BinaryNode
}

func (t *BinaryTree) insert(data int64) *BinaryTree {
	if t.root == nil {
		t.root = &BinaryNode{data: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (n *BinaryNode) insert(data int64) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &BinaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func (t *BinaryTree) print() {
	if t.root == nil {
		return
	}
	fmt.Println("=======================")
	fmt.Println("inOrderPrint")
	t.root.inOrderPrint()

	fmt.Println("=======================")
	fmt.Println("preOrderPrint")
	t.root.preOrderPrint()

	fmt.Println("=======================")
	fmt.Println("postOrderPrint")
	t.root.postOrderPrint()
}

func (n *BinaryNode) inOrderPrint() {
	if n.left != nil {
		n.left.inOrderPrint()
	}
	fmt.Println(n.data)

	if n.right != nil {
		n.right.inOrderPrint()
	}
}

func (n *BinaryNode) preOrderPrint() {
	fmt.Println(n.data)
	if n.left != nil {
		n.left.inOrderPrint()
	}
	if n.right != nil {
		n.right.inOrderPrint()
	}
}

func (n *BinaryNode) postOrderPrint() {
	if n.left != nil {
		n.left.inOrderPrint()
	}
	if n.right != nil {
		n.right.inOrderPrint()
	}
	fmt.Println(n.data)
}

func (t *BinaryTree) contain(data int64) bool {
	if t.root == nil {
		return false
	}
	return t.root.contain(data)
}

func (n *BinaryNode) contain(data int64) bool {
	if n.data == data {
		return true
	}
	if n.data > data {
		if n.left == nil {
			return false
		}
		return n.left.contain(data)
	} else {
		if n.right == nil {
			return false
		}
		return n.right.contain(data)
	}
}
func print(w io.Writer, node *BinaryNode, ns int, ch rune) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.data)
	print(w, node.left, ns+2, 'L')
	print(w, node.right, ns+2, 'R')
}

func BinaryTreePrintAndSearch() {
	tree := &BinaryTree{}
	i := 0
	rand.Seed(time.Now().UnixNano())

	for i < 15 {
		tree.insert(rand.Int63n(50))
		i++
	}
	tree.print()
	print(os.Stdout, tree.root, 0, 'M')

	find := rand.Int63n(50)
	fmt.Println(find, tree.contain(find))
}
