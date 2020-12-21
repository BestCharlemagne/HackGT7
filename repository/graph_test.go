package repository

import "testing"

func TestGraph(t *testing.T) {
	storeGraph := ExampleStores[0].GraphStore()
	print(storeGraph.AdjList[0].StoredItem.Name)
}
