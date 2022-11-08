package dao

import (
	"time"

	"github.com/google/uuid"
)

type Graph interface {
	UpsertLink(link *Link) error
	FindLink(id uuid.UUID) (*Link, error)
	UpsertEdge(edge *Edge) error
	RemoveStaleEdges(fromID uuid.UUID, updatedBefore time.Time) error
	Links(fromID, toID uuid.UUID, retrievedBefore time.Time) (LinkIterator, error)
	Edges(fromID, toID uuid.UUID, updatedBefore time.Time) (EdgeIterator, error)
}
