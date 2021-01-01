package repository

import (
	"container/heap"
)

// An PQItem is something we manage in a priority queue.
type PQItem struct {
	value    *Item // The value of the PQItem; arbitrary.
	priority int   // The priority of the PQItem in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the PQItem in the heap.
}

// A PriorityQueue implements heap.Interface and holds PQItems.
type PriorityQueue []*PQItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	PQItem := x.(*PQItem)
	PQItem.index = n
	*pq = append(*pq, PQItem)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	PQItem := old[n-1]
	old[n-1] = nil    // avoid memory leak
	PQItem.index = -1 // for safety
	*pq = old[0 : n-1]
	return PQItem
}

// update modifies the priority and value of an PQItem in the queue.
func (pq *PriorityQueue) update(PQItem *PQItem, value *Item, priority int) {
	PQItem.value = value
	PQItem.priority = priority
	heap.Fix(pq, PQItem.index)
}

// // This example creates a PriorityQueue with some PQItems, adds and manipulates an PQItem,
// // and then removes the PQItems in priority order.
// func main() {
// 	// Some PQItems and their priorities.
// 	PQItems := map[string]int{
// 		"banana": 3, "apple": 2, "pear": 4,
// 	}

// 	// Create a priority queue, put the PQItems in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(PriorityQueue, len(PQItems))
// 	i := 0
// 	for value, priority := range PQItems {
// 		pq[i] = &PQItem{
// 			value:    value,
// 			priority: priority,
// 			index:    i,
// 		}
// 		i++
// 	}
// 	heap.Init(&pq)

// 	// Insert a new PQItem and then modify its priority.
// 	item := &PQItem{
// 		value:    "orange",
// 		priority: 1,
// 	}
// 	heap.Push(&pq, item)
// 	pq.update(item, item.value, 5)

// 	// Take the PQItems out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*PQItem)
// 		fmt.Printf("%.2d:%s ", item.priority, item.value)
// 	}
// }
