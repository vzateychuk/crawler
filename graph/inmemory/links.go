package inmemory

import (
	"github.com/google/uuid"

	"vez/graph"
)

func (s *GraphInMemory) UpsertLink(link *graph.Link) error {

	// UpsertLink - Since an upsert operation will always modify the graph, the method will acquire
	// a write-lock so that we can apply any modifications in an atomic fashion.
	s.mu.Lock()
	defer s.mu.Unlock()

	// lookup another link with the same URL
	existing := s.linkURLIndex[link.URL]
	if existing == nil {

		// if no existing link found and the link to be upserted does not specify an ID, we generate an ID for the link
		if link.ID == uuid.Nil {
			// Assign new ID
			for {
				link.ID = uuid.New()
				if s.links[link.ID] == nil {
					break
				}
			}
		}

		// Now we make a copy of link to ensure that no code outside of our implementation can modify the graph.
		// Then, we insert the link into the appropriate map structures
		lCopy := new(graph.Link)
		// TODO find out what is going on here
		*lCopy = *link
		s.linkURLIndex[lCopy.URL] = lCopy
		s.links[lCopy.ID] = lCopy
		return nil

	} else {

		// In the case when we've found existing link by URL, silently convert the insert into an update operation
		// while making sure that we always retain the most recent RetrievedAt timestamp:
		link.ID = existing.ID
		origTs := existing.RetrievedAt
		// TODO find out what is going on here
		*existing = *link
		if origTs.After(existing.RetrievedAt) {
			existing.RetrievedAt = origTs
		}
		return nil
	}

}
