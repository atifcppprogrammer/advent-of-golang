package main

import "fmt"

type ListNode struct {
	crate rune
	next  *ListNode
	prev  *ListNode
}

type List struct {
	length int
	tail   *ListNode
}

func (l List) IsEmpty() bool {
	return l.length == 0
}

func (l List) GetLength() int {
	return l.length
}

func (l List) Last() rune {
	if l.length == 0 {
		return '0'
	} else {
		return l.tail.crate
	}
}

func (l List) IsEqual(o List) (bool, string) {
	expected := l.GetLength()
	received := o.GetLength()
	if expected != received {
		return false, fmt.Sprintf(
			"len(this) = %d, len(other) = %d",
			expected,
			received,
		)
	}

	var (
		lCurr = l.tail
		oCurr = o.tail
		index = 0
	)

	for lCurr != nil {
		if lCurr.crate != oCurr.crate {
			return false, fmt.Sprintf(
				"this[%d] = %q, other[%d] = %q",
				l.GetLength()-1-index,
				lCurr.crate,
				o.GetLength()-1-index,
				oCurr.crate,
			)
		}
		lCurr = lCurr.prev
		oCurr = oCurr.prev
		index++
	}
	return true, ""
}

func (l *List) Push(c rune) {
	node := &ListNode{c, nil, nil}
	switch l.length {
	case 0:
		l.tail = node
	default:
		l.tail.next = node
		node.prev = l.tail
		l.tail = l.tail.next
	}
	l.length++
}

func (l *List) Pop() rune {
	var lastCrate rune
	switch l.length {
	case 0:
		break
	case 1:
		lastCrate = l.tail.crate
		l.tail = nil
		l.length--
	default:
		lastCrate = l.tail.crate
		l.tail = l.tail.prev
		l.tail.next = nil
		l.length--
	}
	return lastCrate
}
