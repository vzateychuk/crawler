package inmemory

import "vez/graph"

// the iterator stores a pointer to the in-memory graph, a list of Link models to iterate,
// and an index for keeping track of the iterator's offset within the list.
type linkIterator struct {
	s        *GraphInMemory
	links    []*graph.Link
	curIndex int
}

// Next return true if there is any records left
func (i linkIterator) Next() bool {
	if i.curIndex >= len(i.links) {
		return false
	}
	i.curIndex++
	return true
}

// Link Unless we have already reached the end of the list of links, we
// advance curIndex and return true to indicate that more data is available for
// retrieval via a call to the Link method
func (i *linkIterator) Link() *graph.Link {
	i.s.mu.RLock()
	link := new(graph.Link)
	*link = *i.links[i.curIndex-1]
	i.s.mu.RUnlock()
	return link
}
