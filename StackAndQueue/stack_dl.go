package datastructure

//double linked list

type Element struct {
	Value interface{}
	next, prev *Element 
	stack *Stack
}

type Stack struct {
	root Element
	len int
}

func (s *Stack) Init() *Stack {
	s.root.next = &s.root
	s.root.prev = &l.root 
	l.len = 0
	return l
}

func New() *Stack { return new(Stack).Init() }

func (s *Stack) Len() int { return l.len }

func (s *Stack) insert(e, at *Element) *Element {
	n := at.next
	at.next = e 
	e.prev = at
	e.next = n 
	n.prev = e 
	e.stack = s 
	l.len++ 
	return e
}

/*
func (s *Stack) Push(v interface{}) {
	return s.insert(&Element{Value: v}, &s.root.prev)
}

func (s *Stack) Pop() interface{} {
	if s.root.prev == nil {
		return nil
	}
	
	e := s.root.prev
	e.prev.next = e.next
	e.next.prev = e.prev

	e.next = nil 
	e.prev = nil
	e.stack = nil
	
	return e.Value
}

*/