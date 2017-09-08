package main

import (
	"fmt"
)

type Node struct {
	value       int
	left, right *Node
}

type Tree struct {
	root *Node
	num  int
}

func New() *Tree {
	return new(Tree).Init()
}

func (t *Tree) Init() *Tree {
	t.root = nil
	t.num = 0
	return t
}

func (t *Tree) Insert(v int) {
	n := &Node{v, nil, nil}
	if t.root == nil {
		t.root = n
		return
	}
	cur := t.root
	for cur != nil {
		if cur.value < v {
			if cur.right == nil {
				cur.right = n
				break
			}
			cur = cur.right
		} else {
			if cur.left == nil {
				cur.left = n
				break
			}
			cur = cur.left
		}
	}
	return
}

func (t *Tree) Delete(v int) int {
	//find
	//connect
	//delete
	cur := t.root
	prev := t.root

	if cur.value == v {
		cur = cur.left

		return cur.value
	}
	for cur.value == v {

	}
}

func (t *Tree) Preorder(n *Node) {
	if n == nil {
		return
	}
	fmt.Println(n.value)
	t.Preorder(n.left)
	t.Preorder(n.right)
}

func main() {
	var root *Node
	root = &Node{10, nil, nil}
	root.left = &Node{8, nil, nil}
	root.right = &Node{12, nil, nil}

	fmt.Println(root.value)
	fmt.Println(root.left.value)
	fmt.Println(root.right.value)

	t := New()
	fmt.Println(t.root)
	t.Insert(10)
	t.Insert(8)
	t.Insert(12)
	t.Preorder(t.root)
}
