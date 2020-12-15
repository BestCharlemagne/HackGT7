type StoreGraph struct {
	AdjList map[Vertex]VertexDistance `json:"adjList"`
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
func (s Store) GraphStore() GraphedStore {
	graphedStore := GraphvedStore{}
	if len(s.Path) > 0 && len(s.Path[0]) > 0 {
		graphedStore.GraphedStore = s
		height := len(s.Path)
		width := len(s.Path[0])

		queue := StoreQueue{}
		queue.push()

		i := 0
		j := 0

		for i < height && j < width {
			currRune := s.Path[i][j]
			currType := currRune.getItemType()
		}
	}

	return graphedStore
}
