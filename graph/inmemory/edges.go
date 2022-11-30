package inmemory

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

	"vez/graph"
)

func (s *GraphInMemory) UpsertEdge(edge *graph.Edge) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	//  verify that the source and destination links for the edge actually exist
	_, srcExists := s.links[edge.Src]
	_, dstExists := s.links[edge.Dst]
	if !srcExists || !dstExists {
		return xerrors.Errorf("upsert edge: %w", graph.ErrUnknownEdgeLinks)
	}

	// Scan  the set of edges that originate from the specified source link and check whether we can find
	//an existing edge to the same destination. If that happens to be the case, we simply update the
	//entry's UpdatedAt field and copy its contents back to the provided edge pointer.
	//This ensures that the entry value that's provided by the caller has both its ID and UpdatedAt synced
	//with the values contained in the store
	now := time.Now()
	for _, edgeID := range s.linkEdgeMap[edge.Src] {
		existingEdge := s.edges[edgeID]
		if existingEdge.Src == edge.Src && existingEdge.Dst == edge.Dst {
			existingEdge.UpdatedAt = now
			*edge = *existingEdge
			return nil
		}
	}
	// If the preceding loop does not produce a match, we create and insert a new edge to the store.
	for {
		edge.ID = uuid.New()
		if s.edges[edge.ID] == nil {
			break
		}
	}

	edge.UpdatedAt = time.Now()
	eCopy := new(graph.Edge)
	*eCopy = *edge
	s.edges[eCopy.ID] = eCopy
	// Append the edge ID to the list of edges originating from the edge's source link.
	s.linkEdgeMap[edge.Src] = append(s.linkEdgeMap[edge.Src], eCopy.ID)
	return nil

}

func (s *GraphInMemory) RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var newEdgeList edgeList
	for _, edgeID := range s.linkEdgeMap[fromID] {
		edge := s.edges[edgeID]
		if edge.UpdatedAt.Before(updatedBefore) {
			delete(s.edges, edgeID)
		} else {
			newEdgeList = append(newEdgeList, edgeID)
		}
	}
	s.linkEdgeMap[fromID] = newEdgeList
	return nil
}
