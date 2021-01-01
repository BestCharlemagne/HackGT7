package repository

import "container/heap"

//StoreGraph contains all data pertaining to the graph of stores
//Both the graph and the store have an Items structure. The Items in the Store is a
//simple list of Items, while the Items in this struct is a map that allows quick access to a specific item by ID.
type StoreGraph struct {
	GraphedStore *Store
	Items        map[int]*Item `json:"items"`
	AdjList      []Vertex      `json:"adjList"`
	StartVertex  *Vertex       `json:"start"`
	EndVertex    *Vertex       `json:"end"`
}

//Vertex contains the data at the vertex as well as distances to all neighbor vertices
type Vertex struct {
	StoredItem *Item            `json:"item"`
	Neighbors  []VertexDistance `json:"neighbors"`
}

//VertexDistance contains the distance to a specific vertex
type VertexDistance struct {
	DestinationVertex *Vertex `json:"item"`
	Distance          int     `json:"distance"`
}

//Iterates over store path to create a complete graph of it.
func (s *Store) GraphStore() StoreGraph {
	graph := StoreGraph{}
	if len(s.Path) > 0 && len(s.Path[0]) > 0 {
		graph.GraphedStore = s
		height := len(s.Path)
		width := len(s.Path[0])

		graph.AdjList = make([]Vertex, height*width)
		var j int //Row
		i := 0    //Column

		graph.Items = make(map[int]*Item, len(s.Items))

		memo := make(map[int]*Item, len(s.Items))

		index := 0

		for j = 0; j < len(s.Path); j++ {
			for i = 0; i < len(s.Path[j]); i++ {
				if graph.GraphedStore.Path[j][i].getItemType() != Wall {
					neighborVertex := Vertex{StoredItem: graph.getItem(memo, &index, j, i)}
					graph.pushNeighbors(&neighborVertex, j, i-1)
					graph.pushNeighbors(&neighborVertex, j, i+1)
					graph.pushNeighbors(&neighborVertex, j-1, i)
					graph.pushNeighbors(&neighborVertex, j+1, i)
					graph.AdjList[getIndex(&s.Path, j, i)] = neighborVertex

					if neighborVertex.StoredItem.Type == Entrance {
						graph.StartVertex = &neighborVertex
					} else if neighborVertex.StoredItem.Type == Exit {
						graph.EndVertex = &neighborVertex
					}
				}
			}
		}
	}

	return graph
}

//getIndex converts a row and column to a single number to be accessed easily from list
func getIndex(path *PathArray, row int, column int) int {
	return row*len((*path)[row]) + column
}

//getItem expects a reference to an index that will be able to consistently store a value.
//Caches all of the incremented values to lazily build the graph and retrieve the item simultaneously.
//Also uses memoization to quickly access already cached Items that have not yet been used
func (graph *StoreGraph) getItem(memo map[int]*Item, index *int, row int, column int) *Item {
	var itemAtIndex *Item = memo[getIndex(&graph.GraphedStore.Path, row, column)]
	for itemAtIndex == nil && *index < len(graph.GraphedStore.Items) {
		if graph.GraphedStore.Items[*index].Row == row && graph.GraphedStore.Items[*index].Column == column {
			itemAtIndex = &graph.GraphedStore.Items[*index]
		}
		graph.Items[graph.GraphedStore.Items[*index].ID] = &graph.GraphedStore.Items[*index]
		cachedItem := &graph.GraphedStore.Items[*index]
		memo[getIndex(&graph.GraphedStore.Path, cachedItem.Row, cachedItem.Column)] = cachedItem
		*index++
	}
	if itemAtIndex == nil {
		itemAtIndex = &Item{Type: graph.GraphedStore.Path[row][column].getItemType(), Row: row, Column: column}
	}
	return itemAtIndex
}

//pushNeighbors adds all neighbors of a vertex to the graph
func (graph *StoreGraph) pushNeighbors(neighborVertex *Vertex, row int, column int) {
	if row >= 0 && row < len(graph.GraphedStore.Path) && column >= 0 &&
		column < len(graph.GraphedStore.Path[row]) && graph.GraphedStore.Path[row][column].getItemType() != Wall {
		vertex := &graph.AdjList[getIndex(&graph.GraphedStore.Path, row, column)]
		distance := abs(neighborVertex.StoredItem.Row-row) + abs(neighborVertex.StoredItem.Column-column)
		neighborVertex.Neighbors = append(neighborVertex.Neighbors, VertexDistance{DestinationVertex: vertex, Distance: distance})
	}
}

func (graph *StoreGraph) findPathBetweenItems(items []*Item) []*Item {
	path := []*Item{}
	itemMap := make(map[*Item]bool, len(items))
	for _, value := range items {
		itemMap[value] = true
	}
	currentItem := graph.StartVertex.StoredItem
	for currentItem != nil && len(itemMap) != 0 {
		frontier := &PriorityQueue{}
		heap.Init(frontier)
		heap.Push(frontier, &PQItem{value: currentItem, priority: 0})
		cameFrom := map[int]int{}
		costSoFar := map[int]int{}

		cameFrom[graph.itemToInt(graph.StartVertex.StoredItem)] = 0
		costSoFar[graph.itemToInt(graph.StartVertex.StoredItem)] = 0

		var foundItem *Item

		for frontier.Len() != 0 && foundItem == nil {
			curr := heap.Pop(frontier).(*PQItem).value

			if itemMap[curr] {
				foundItem = curr
			} else {
				for _, next := range graph.AdjList[getIndex(&graph.GraphedStore.Path, curr.Row, curr.Column)].Neighbors {
					newCost := costSoFar[getIndex(&graph.GraphedStore.Path, curr.Row, curr.Column)] + 1
					_, present := costSoFar[graph.itemToInt(next.DestinationVertex.StoredItem)]
					if !present || newCost < costSoFar[graph.itemToInt(next.DestinationVertex.StoredItem)] {
						priority := newCost + heuristic(items, next)
						heap.Push(frontier, &PQItem{value: next.DestinationVertex.StoredItem, priority: priority})
						cameFrom[graph.itemToInt(next.DestinationVertex.StoredItem)] = graph.itemToInt(curr)
					}
				}
			}

		}

		delete(itemMap, foundItem)
		path = append(path, foundItem)
		currentItem = foundItem
	}

	return path
}

func heuristic(items []*Item, item VertexDistance) int {
	return 1
}

func (graph *StoreGraph) itemToInt(i *Item) int {
	return getIndex(&graph.GraphedStore.Path, i.Row, i.Column)
}

//abs value of an integer
func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
