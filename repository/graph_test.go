package repository

import (
	"fmt"
	"testing"
)

var tinyExampleStore Store = Store{
	ID:      1,
	Title:   "Joe's",
	ZipCode: 1,
	Items:   []Item{{0, "Banana", 0, 0, 1}},
	Path:    [][]pathRune{{'+', '-'}, {'_', '+'}},
}

func TestTinyStore(t *testing.T) {
	storeGraph := tinyExampleStore.GraphStore()

	correctStart := checkProperLocation(storeGraph.StartVertex, 0, 1)
	if correctStart != "" {
		t.Fatal(correctStart)
	}

	correctEnd := checkProperLocation(storeGraph.EndVertex, 1, 0)
	if correctEnd != "" {
		t.Fatal(correctEnd)
	}

	adjacencyList := []Vertex{
		{StoredItem: &tinyExampleStore.Items[0]},
		{StoredItem: &Item{Type: Entrance, Row: 0, Column: 1}, Neighbors: []VertexDistance{}},
		{StoredItem: &Item{Type: Exit, Row: 1, Column: 0}, Neighbors: []VertexDistance{}},
		{StoredItem: &Item{Type: Empty, Row: 1, Column: 1}, Neighbors: []VertexDistance{}}}

	adjacencyList[0].Neighbors = []VertexDistance{{DestinationVertex: &adjacencyList[1], Distance: 1}, {DestinationVertex: &adjacencyList[2], Distance: 1}}
	adjacencyList[1].Neighbors = []VertexDistance{{DestinationVertex: &adjacencyList[0], Distance: 1}, {DestinationVertex: &adjacencyList[3], Distance: 1}}
	adjacencyList[2].Neighbors = []VertexDistance{{DestinationVertex: &adjacencyList[0], Distance: 1}, {DestinationVertex: &adjacencyList[3], Distance: 1}}
	adjacencyList[3].Neighbors = []VertexDistance{{DestinationVertex: &adjacencyList[1], Distance: 1}, {DestinationVertex: &adjacencyList[2], Distance: 1}}

	compareAdjacencyLists(adjacencyList, storeGraph.AdjList)
}

func compareAdjacencyLists(a1 []Vertex, a2 []Vertex) string {
	if len(a1) != len(a2) {
		return "Adjacency lists have different sizes"
	}

	for index, value := range a1 {
		v1 := value
		v2 := a2[index]

		sameItem := *v1.StoredItem == *v2.StoredItem

		if !sameItem {
			return fmt.Sprintf(`Neighbors at index %v have diverged. They have different items.`, index)
		}

		if len(v1.Neighbors) != len(v2.Neighbors) {
			return fmt.Sprintf(`Neighbors at index %v have different lengths`, index)
		}

		for i, v := range v1.Neighbors {
			n1 := v
			n2 := v2.Neighbors[i]
			sameDestination := n1.DestinationVertex == n2.DestinationVertex
			sameDistance := n1.Distance == n2.Distance

			if !sameDestination && !sameDistance {
				return fmt.Sprintf(`Vertices at index %v have diverged. They have different neighbors at index %v.`, index, i)
			}
		}
	}
	return ""
}

func checkProperLocation(v *Vertex, r int, c int) string {
	if v.StoredItem.Row != r || v.StoredItem.Column != c {
		return fmt.Sprintf(`Start Vertex was supposed to be at row: %v and column: %v but was found at row: %v and column: %v`, r, c, v.StoredItem.Row, v.StoredItem.Column)
	}
	return ""
}
