package main

import(
	"fmt"
)


type Node struct {
	value interface{}
	left, right *Node
	balfactor int8
}


func updateBalanceFactor(A *Node) (r int8) {
	if A.left != nil && A.right != nil {
		r = A.left.balfactor - A.right.balfactor
	} else if A.left != nil && A.right == nil {
		r = 1
	} else { // A.left == nil && A.right != nil {
		r = -1
	}
	return
}

func main() {
	A := &Node{5, nil, nil, 0}

/*
		   A
	    /    \
       B      nil
*/
	B := &Node{3, nil, nil, 0}
	A.left = B
	A.balfactor = updateBalanceFactor(A)
	fmt.Println("A :", A)
	fmt.Println("B :", B)


/*
		   A
	    /    \
       B      C
*/
	C:= &Node{6, nil, nil, 0}	
	A.right = C
	A.balfactor = updateBalanceFactor(A)
	fmt.Println("A: ", A)
}