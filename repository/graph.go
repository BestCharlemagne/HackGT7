type StoreGraph struct {
	Vertices []Vertex                    `json:"vertices"`
	AdjList  map[Vertex][]VertexDistance `json:"adjList"`
}

type Vertex struct {
	StoredItem Item `json:"item"`
}

type VertexDistance struct {
	DestinationItem Item `json:"item"`
	Distance        int  `json:"distance"`
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
		size := height * width
		visitedMap := map[Item]bool
		queue := StoreQueue{}
		queue.push(Item{Type: s.Path[0][0].getItemType(), Row: 0, Column: 0})

		for len(visitedMap) < size {
			currItem = graphedStore.pop()
			if !currItem[i] {
				visitedMap[i] = true
				//Push left neighbor
				if currItem.Column > 0 {
					pushItem(currItem.Row, currItem.Column - 1)
				}
				//Push top neighbor
				if currItem.Row > 0 {
					pushItem(currItem.Row - 1, currItem.Column)
				}
				//Push right neighbor
				if currItem.Column < width {
					pushItem(currItem.Row, CurrItem.Column + 1)
				}
				//Push bottom neighbor
				if currItem.Row < height {
					pushItem(currItem.Row + 1, CurrItem.Column)
				}
			}
		}
	}

	func (curr Item) pushItem(row int, column int) {
		//TODO: need to check if a product is at this location
		item := Item{Type: s.Path[row][column]}, Row: row, Column: column}
		if !visitedMap[item] {
			queue.push(item)
			append(graphedStore.Graph.AdjList[Item], VertexDistance{item, abs(curr.Row - row) + abs(curr.Column - column)}
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
