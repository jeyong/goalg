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

func updateBF() {
//관찰 
/*
언제 bf가 바뀌는가?
bf가 바뀌는 것들의 특징이 있는가?
그 rule을 찾아내면 쉽게 프로그래밍이 가능할텐데!

알고리즘을 외우기 보다는
현상을 관찰하고 거기에서 특징을 찾아내고 그것을 rule로 만들어 내는 방식을 연습해야한다!

				   A
			   B       C
			D              E


에서 'F'가 추가되는 경우
				   A
			   B       C
			D     F        E
B와 A의 bf값이 변경되게 된다. 다른 node의 bf는 변화가 없다.

다른 경우는 어떻게 될까?

*/

	for(x := getParent(z); x != nil; x = getParent(z)){
		z = x
	}
}

func main() {
/*
		   A
	    /    \
     nil     nil
*/

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


