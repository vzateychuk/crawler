package inmemory

import (
	"sync"

	"github.com/google/uuid"

	"vez/graph"
)

type edgeList []uuid.UUID

type GraphInMemory struct {
	// ensure that our implementation is safe for concurrent access
	mu sync.RWMutex

	//  links/edges maintain the set of Link and Edge models that have been inserted into the graph.
	links map[uuid.UUID]*graph.Link
	edges map[uuid.UUID]*graph.Edge

	// maintains an auxiliary map (linkURLIndex) where keys are the URLs that are added to the graph and values are pointers to link models
	linkURLIndex map[string]*graph.Link
	// This map associates link IDs with a slice of edge IDs that correspond to the edges originating from it
	linkEdgeMap map[uuid.UUID]edgeList
}
