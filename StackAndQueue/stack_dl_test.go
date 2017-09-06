package datastructure

import "testing"

//Testing은 실제 Test의 경우 함수 TestXxxxx로 시작하며
//실제 check하는 함수는 checkXxxx로 시작함.
//최대한 Test에는 실제 사용 시나리오만 넣도록 한다.

// Test Push with unitialized Stack
func TestZeroStack(t *testing.T) {
	var s1 = new(Stack)
	s1.Push(1)
	checkStack(t, s1, []interface{}{1})
}

func TestPop(t *testint.T) {
	var s Stack 
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Pop()
	checkStack(t, &s, []interface{}{1, 2, 3})
}

func TestPush(t *testing.T) {

}

func TestStack(t *testing.T) {
	
}

	
func checkStackLen(t *testing.T, s *Stack, es []interface{}) {
	//check len
	//
}

func checkStackPointers(t *testing.T, s *Stack, es []*Element) {

}

func checkStack(t *testing.T, s *Stack, es []interface{}) {

}

func checkStackValue(t *testint.T, v1 interface{}, v2 interface{}){
	if v1 != v2 {
		t.Errorf("%v is not same as %v", v1, v2)
	}
}