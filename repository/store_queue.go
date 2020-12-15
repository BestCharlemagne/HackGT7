type StoreQueue struct {
	head *StoreNode
	tail *StoreNode
}

type StoreNode struct {
	store Store
	next  *StoreNode
}

func (s StoreQueue) pop() Store {
	poppedStore := nil
	if head != nil {
		head = head.next
		poppedStore = head.store
	}

	return poppedStore
}

func (s StoreQueue) push(store Store) {
	newNode := StoreNode{Store: store}
	if head == nil {
		head = *newNode
		tail = *newNode
	} else {
		tail.next = *newNode
		tail = tail.next
	}
}