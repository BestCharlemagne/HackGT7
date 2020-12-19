package repository

//ItemQueue is the structure to hold the head and tail of a queue
type ItemQueue struct {
	head *ItemNode
	tail *ItemNode
}

//ItemNode is a singly linked list item
type ItemNode struct {
	item Item
	next *ItemNode
}

//pop returns and removes the next item of the queue
func (s ItemQueue) pop() Item {
	var poppedItem Item
	if s.head != nil {
		s.head = s.head.next
		poppedItem = s.head.item
	}

	return poppedItem
}

//push adds a new item to the queue
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
