package graph

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID          uuid.UUID
	URL         string
	RetrievedAt time.Time
}
