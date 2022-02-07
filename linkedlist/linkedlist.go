package linkedlist

type Element struct {
	Value string
	next  *Element
	prev  *Element
}

func (e Element) Next() *Element {
	return e.next
}

func (e Element) Prev() *Element {
	return e.prev
}

type List struct {
	head *Element
	tail *Element
}

func New() *List {
	return &List{}
}

func (l List) Head() *Element {
	return l.head
}

func (l *List) Queue(value string) *Element {
	elem := Element{
		Value: value,
		prev:  l.tail,
	}
	if l.tail != nil {
		l.tail.next = &elem
	}

	if l.head == nil {
		l.head = &elem
	}
	l.tail = &elem

	return &elem
}

func (l *List) Dequeue() *Element {
	elem := l.head
	if elem.next != nil {
		elem.next.prev = nil
	}
	if l.head != nil {
		l.head = l.head.next
	}

	return elem
}
