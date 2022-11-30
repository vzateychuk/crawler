package inmemory

import "vez/graph"

// the iterator stores a pointer to the in-memory graph, a list of Edge models to iterate,
// and an index for keeping track of the iterator's offset within the list.
type edgeIterator struct {
	s        *GraphInMemory
	edges    []*graph.Edge
	curIndex int
}

// Next return true if there is any records left
func (i edgeIterator) Next() bool {
	if i.curIndex >= len(i.edges) {
		return false
	}
	i.curIndex++
	return true
}

// Edge Unless we have already reached the end of the list of edges, we
// advance curIndex and return true to indicate that more data is available for
// retrieval via a call to the Edge method
func (i *edgeIterator) Edge() *graph.Edge {
	i.s.mu.RLock()
	edge := new(graph.Edge)
	*edge = *i.edges[i.curIndex-1]
	i.s.mu.RUnlock()
	return edge
}
