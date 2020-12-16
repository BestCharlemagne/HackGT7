package repository

type ItemQueue struct {
	head *ItemNode
	tail *ItemNode
}

type ItemNode struct {
	item Item
	next *ItemNode
}

func (s ItemQueue) pop() Item {
	var poppedItem Item
	if s.head != nil {
		s.head = s.head.next
		poppedItem = s.head.item
	}

	return poppedItem
}

func (s ItemQueue) push(item Item) {
	newNode := ItemNode{item: item}
	if s.head == nil {
		s.head = &newNode
		s.tail = &newNode
	} else {
		s.tail.next = &newNode
		s.tail = s.tail.next
	}
}
