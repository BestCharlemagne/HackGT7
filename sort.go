package main

type distanceContainer struct {
	distance int
	stored   *Store
}

//DistanceLinkedListNode is a simple struct for a doubly linked list.
type distanceLinkedListNode struct {
	data distanceContainer
	next *distanceLinkedListNode
	prev *distanceLinkedListNode
}

//DistanceLinkedList creates a linked list that is used to sort items by relation to target.
//Infinite size of size set to -1.
type DistanceLinkedList struct {
	front   *distanceLinkedListNode
	back    *distanceLinkedListNode
	target  int
	size    int
	maxSize int
}

//Add adds an item to a DistanceLinkedList at the proper location.
//Only keeps size closest elements.
func (d *DistanceLinkedList) Add(store *Store) {
	distance := abs(store.ZipCode - d.target)
	node := &distanceLinkedListNode{data: distanceContainer{distance: distance, stored: store}, prev: d.back}
	if d.front == nil {
		d.front = node
		d.back = node
	} else {
		if distance >= d.back.data.distance {
			if d.size == d.maxSize {
				return
			}
			d.back.next = node
			d.back = d.back.next
		} else {
			if distance < d.front.data.distance {
				d.front.prev = node
				d.front = d.front.prev
			} else {
				curr := d.front
				for i := 0; i < d.size; i++ {
					if curr != nil && distance < curr.data.distance {
						curr.prev.next = node
						curr.prev = node
					}
					curr = curr.next
				}
			}
			if d.size > d.maxSize {
				d.back = d.back.prev
				d.back.next = nil
			}
		}
	}
	d.size++
}

func (d DistanceLinkedList) toArray() []Store {
	array := make([]Store, d.size)
	curr := d.front
	for i := 0; i < d.size; i++ {
		array[i] = *curr.data.stored
		curr = curr.next
	}
	return array
}

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
