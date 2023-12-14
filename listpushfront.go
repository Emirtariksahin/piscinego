package piscine

func ListPushFront(l *List, data interface{}) {
	newnode := &NodeL{Data: data, Next: l.Tail}

	if l.Head == nil {
		l.Head = newnode
		l.Tail = newnode
	} else {

		l.Head = newnode
		l.Tail = newnode
	}
}

/*n := &NodeL{Data: data}
if l.Head == nil {
	l.Head = n
	l.Tail = n
} else {
	l.Tail.Next = n
	l.Tail = n
}*/
