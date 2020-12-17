package repository

type StoreGraph struct {
	Items       map[int]Item `json:"items"`
	AdjList     []Vertex     `json:"adjList"`
	StartVertex *Vertex      `json:"start"`
	EndVertex   *Vertex      `json:"end"`
}

type Vertex struct {
	StoredItem Item     `json:"item"`
	Neighbors  []Vertex `json:"neighbors"`
}

type VertexDistance struct {
	DestinationVertex Vertex `json:"item"`
	Distance          int    `json:"distance"`
}

type GraphedStore struct {
	GraphedStore Store      `json:"store"`
	Graph        StoreGraph `json:"graph"`
}

//GraphStore will only be successful if path is connected
//Uses breadth-first-search to combine path and items into graph
//TODO: Doesn't need to use breadth-first-search, can just iteratate over all entries
func (s Store) GraphStore() GraphedStore {
	graphedStore := GraphedStore{}
	if len(s.Path) > 0 && len(s.Path[0]) > 0 {
		graphedStore.GraphedStore = s
		height := len(s.Path)
		width := len(s.Path[0])

		graphedStore.Graph.AdjList = make([]Vertex, height*width)
		var j int //Row
		i := 0    //Column

		getIndex := func(row int, column int) int {
			return row*len(s.Path[row]) + column
		}

		index := 0
		getItem := func(row int, column int) Item {
			var itemAtIndex *Item
			for itemAtIndex != nil && index < len(s.Items) {
				if s.Items[index].Row == row && s.Items[index].Column == column {
					itemAtIndex = &s.Items[index]
				}
				graphedStore.Graph.Items[s.Items[index].ID] = s.Items[index]
				index++
			}
			if itemAtIndex == nil {
				itemAtIndex = &Item{Type: s.Path[row][column].getItemType(), Row: row, Column: column}
			}
			return *itemAtIndex
		}

		//TODO: Need to incorporate distance
		pushItem := func(neighborVertex Vertex, row int, column int) {
			if row >= 0 && row < len(s.Path) && column >= 0 && column < len(s.Path[row]) && s.Path[row][column].getItemType() != Wall {
				neighborVertex.Neighbors = append(neighborVertex.Neighbors, graphedStore.Graph.AdjList[getIndex(row, column)])
			}
		}

		for j = 0; j < len(s.Path); j++ {
			for i < len(s.Path[j]) {
				neighborVertex := graphedStore.Graph.AdjList[getIndex(i, j)]
				pushItem(neighborVertex, i, j-1)
				pushItem(neighborVertex, i, j+1)
				pushItem(neighborVertex, i-1, j)
				pushItem(neighborVertex, i+1, j)
				graphedStore.Graph.AdjList[getIndex(i, j)] = Vertex{StoredItem: getItem(i, j)}
			}
		}
	}

	return graphedStore
}

//abs value of an integer
func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}
