package linked_list

type node struct {
	Data int
	next *node
}

type linkedList struct {
	Length int
	Head   *node
	Tail   *node
}

func (l *linkedList) Append(n *node) {
	if l.Head == nil {
		l.Head = n
	}
	if l.Tail != nil {
		l.Tail.next = n
	}
	l.Tail = n
	l.Length++
}

func (l *linkedList) Unshift(n *node) {
	if l.Head == nil {
		l.Head = n
	}
	if l.Tail != nil {
		l.Tail.next = n
	}
	l.Tail = n
	l.Length++
}
