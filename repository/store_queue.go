type ItemQueue struct {
	head *ItemNode
	tail *ItemNode
}

type ItemNode struct {
	item Item
	next *ItemNode
}

func (s StoreQueue) pop() Item {
	poppedItem := nil
	if head != nil {
		head = head.next
		poppedStore = head.item
	}

	return poppedItem
}

func (s ItemQueue) push(item Item) {
	newNode := ItemNode{item: item}
	if head == nil {
		head = *newNode
		tail = *newNode
	} else {
		tail.next = *newNode
		tail = tail.next
	}
}