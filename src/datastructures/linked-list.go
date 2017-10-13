package datastructures


type LinkedList struct {
	head *listNode
	last *listNode
	size int
}
type listNode struct {
	val int
	next *listNode
	prev *listNode
}
func (l *LinkedList) Push(value int) int {
	return l.Add(value)
}
func (l *LinkedList) Pop() (int, bool) {
	if l.size == 0 {
		return 0, false
	}
	val := l.head.val
	if l.size == 1 {
		l.head = nil
		l.last = nil
	} else {
		l.head = l.head.next
	}
	l.size--
	return val, true
}
func (l *LinkedList) Add(value int) int {
	newNode := listNode{val:value}
	if l.head == nil {
		l.head = &newNode
		l.last = l.head
	} else {
		l.last.next = &newNode
		newNode.prev = l.last
		l.last = &newNode
	}
	l.size++
	return value
}
func (l *LinkedList) Remove(value int) bool {
	if l.size == 0 {
		return false
	} else if l.head.val == value {
		if l.size == 1 {
			l.head = nil
			l.last = nil
		} else {
			l.head.next.prev = nil
			l.head = l.head.next
			l.last = l.head
		}
		l.size--
		return true
	}
	iter := l.head.next
	for iter != nil {
		if iter.val == value {
			if iter.next != nil {
				iter.prev.next = iter.next
				iter.next.prev = iter.prev
			} else {
				iter.prev.next = nil
				l.last = iter.prev
			}
			l.size--
			return true
		}
		iter = iter.next
	}
	return false
}
func (l *LinkedList) Contains(value int) bool {
	if l.size == 0 {
		return false
	}
	iter := l.head
	for iter != nil {
		if iter.val == value {
			return true
		}
		iter = iter.next
	}
	return false
}