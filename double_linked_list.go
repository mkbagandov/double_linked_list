package double_linked_list

type node struct {
	val  map[string]string
	next *node
	prev *node
}

type DoubleLinkedList struct {
	tail *node
	head *node
}

func (l *DoubleLinkedList) Append(val map[string]string) {
	if l.head == nil {
		l.head = &node{val: val}
		l.tail = l.head
		return
	}
	l.tail.next = &node{val: val, prev: l.tail}
	l.tail = l.tail.next
}

func (l *DoubleLinkedList) Prepend(val map[string]string) {
	if l.head == nil {
		l.head = &node{val: val}
		l.tail = l.head
		return
	}
	l.head.prev = &node{val: val, next: l.head}
	l.head = l.head.prev
}

func (l *DoubleLinkedList) Print() *[]map[string]string {
	if l.head != nil {
		result := make([]map[string]string, 1)
		for {
			result = append(result, l.head.val)
			if l.head.next == nil {
				break
			}
			l.head = l.head.next
		}
		return &result
	}
	return nil
}

func (l *DoubleLinkedList) PrintReverse() *[]map[string]string {
	if l.tail != nil {
		result := make([]map[string]string, 1)
		for {
			result = append(result, l.tail.val)
			if l.tail.prev == nil {
				break
			}
			l.tail = l.tail.prev
		}
		return &result
	}
	return nil
}
