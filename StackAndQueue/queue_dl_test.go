package datastructure

import "testing"

//Testing은 실제 Test의 경우 함수 TestXxxxx로 시작하며
//실제 check하는 함수는 checkXxxx로 시작함.
//최대한 Test에는 실제 사용 시나리오만 넣도록 한다.

// Test Push with unitialized Stack
func TestZeroQueue(t *testing.T) {
	var q1 = new(Queue)
	q1.Push(1)
	checkStack(t, q1, []interface{}{1})
}

func TestPop(t *testint.T) {
	var q Queue 
	q.Push(1)
	q.Push(2)
	q.Push(3)

	q.Pop()
	checkStack(t, &q, []interface{}{1, 2, 3})
}

func TestPush(t *testing.T) {

}

func TestQueue(t *testing.T) {
	
}

	
func checkQueueLen(t *testing.T, q *Queue, es []interface{}) {
	//check len
	//
}

func checkQueuePointers(t *testing.T, q *Queue, es []*Element) {

}

func checkQueue(t *testing.T, q *Queue, es []interface{}) {

}

func checkQueueValue(t *testint.T, v1 interface{}, v2 interface{}){
	if v1 != v2 {
		t.Errorf("%v is not same as %v", v1, v2)
	}
}