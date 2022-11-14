package graph

/*
Here introduced generic Iterator.
To iterate a list of edges or links, we must obtain an iterator from the graph and run our business logic within a for loop:
('linkIt' is a link iterator)

	for linkIt.Next() {
		link := linkIt.Link()
		// Do something with link...
	}

	if err := linkIt.Error(); err != nil {
		// Handle error...
	}

Calls to linkIt.Next() will return false when the following occurs:
- We have iterated all the available links;
- An error occurs (for example, we lost connection to the database);

As a result, we don't need to check whether an error occurred inside the
loop – we only need to check once after exiting the for loop. This pattern
yields cleaner-looking code and is actually used in various places within
the Go standard library, such as the Scanner type from the bufio package.
*/
type Iterator interface {

	// Next advances the iterator. If no more items are available or an error occurs, calls to Next() return false.
	Next() bool
	// Error returns the last error encountered by the iterator.
	Error() error
	// Close releases any resources associated with an iterator.
	Close() error
}

/*
Both of the followed interfaces define a getter method for retrieving
the Link or Edge instance that the iterator is currently pointing at. The
common logic between the two iterators has been extracted into a
separate interface called Iterator, which both of the interfaces embed.
*/
type LinkIterator interface {
	Iterator[Link]

	// Link returns the currently fetched link object.
	Link() *Link
}

type EdgeIterator interface {
	Iterator[Edge]

	// Edge returns the currently fetched edge objects.
	Edge() *Edge
}
