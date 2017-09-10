package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	left, right *Node 
}


func main() {
	A := new(Node) // literal init : &Node{1, nil, nil}
	A.value = 2
	A.left, A.right = nil, nil 
	fmt.Println(A)

/*
		   A
	    /    \
       B      C
*/
	B := &Node{3, nil, nil}
	A.left = B

	C := &Node{4, nil, nil}
	A.right = C

	fmt.Println(A)
	fmt.Println(A.left)
	fmt.Println(A.right)
}