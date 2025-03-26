package double_linked_list

import "fmt"

type Node struct {
	val  string
	next *Node
	prev *Node
}

type DoubleLinkedList struct {
	tail *Node
	head *Node
}

func (l *DoubleLinkedList) Append(val string) {
	if l.head == nil {
		l.head = &Node{val: val}
		l.tail = l.head
		return
	}
	l.tail.next = &Node{val: val, prev: l.tail}
	l.tail = l.tail.next
}

func (l *DoubleLinkedList) Prepend(val string) {
	if l.head == nil {
		l.head = &Node{val: val}
		l.tail = l.head
		return
	}
	l.head.prev = &Node{val: val, next: l.head}
	l.head = l.head.prev
}

func (l *DoubleLinkedList) Print() {
	if l.head != nil {
		for {
			fmt.Println(l.head.val)
			if l.head.next == nil {
				break
			}
			l.head = l.head.next
		}
	}
}

func (l *DoubleLinkedList) PrintReverse() {
	if l.tail != nil {
		for {
			fmt.Println(l.tail.val)
			if l.tail.prev == nil {
				break
			}
			l.tail = l.tail.prev
		}
	}
}
