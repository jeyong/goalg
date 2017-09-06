package datastructure

//double linked list

type Element struct {
	Value interface{}
	next, prev *Element 
	queue *Stack
}

type Queue struct {
	root Element
	len int
}

func (q *Queue) Init() *Queue {
	q.root.next = &q.root
	q.root.prev = &l.root 
	l.len = 0
	return l
}

func New() *Queue { return new(Queue).Init() }

func (q *Queue) Len() int { return l.len }

func (q *Queue) insert(e, at *Element) *Element {
	n := at.next
	at.next = e 
	e.prev = at
	e.next = n 
	n.prev = e 
	e.queue = q 
	l.len++ 
	return e
}

/*
func (q *Queue) Push(v interface{}) {
	return q.insert(&Element{Value: v}, &q.root.prev)
}

func (q *Queue) Pop() interface{} {
	if q.root.next == nil {
		return nil
	}

	e := q.root.next
	e.prev.next = e.next
	e.next.prev = e.prev

	e.next = nil 
	e.prev = nil
	e.queue = nil
	
	return e.Value
}

*/