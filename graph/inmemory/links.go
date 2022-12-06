package inmemory

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/xerrors"

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

func (s *GraphInMemory) FindLink(id uuid.UUID) (*graph.Link, error) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	link := s.links[id]
	if link == nil {
		return nil, xerrors.Errorf("find link: %w", graph.ErrNotFound)
	}

	// Since we want to ensure that no external code can modify the graph's
	//contents without invoking the UpsertLink method,
	//the FindLink implementation always returns a copy of the link that is
	//stored in the graph.
	lCopy := new(graph.Link)
	*lCopy = *link
	return lCopy, nil
}

func (s *GraphInMemory) Links(fromID, toID uuid.UUID, retrievedBefore time.Time) (graph.LinkIterator, error) {
	from, to := fromID.String(), toID.String()

	s.mu.RLock()
	defer s.mu.RUnlock()

	// iterate all the links in the graph, searching for the ones that belong to
	// the [fromID, toID) partition range and whose RetrievedAt value is less than
	// the specified retrievedBefore
	var list []*graph.Link
	for linkId, link := range s.links {
		if id := linkId.String(); id >= from && id < to && link.RetrievedAt.Before(retrievedBefore) {
			list = append(list, link)
		}
	}

	// create a new linkIterator instance and return it to the user
	iterator := &linkIterator{s: s, links: list}
	return iterator, nil
}
