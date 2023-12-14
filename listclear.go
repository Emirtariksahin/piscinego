package piscine

func ListClear(l *List) {
	if l.Head == nil {
		return
	}
	l.Head = nil
	l.Tail = nil
}
