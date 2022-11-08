# DAO - Data Access Layer

the link graph access layer must support the following set of operations:

1. Insert a link into the graph or update an existing link when the  crawler 
   discovers that its content has changed.
2. Look up a link by its ID.
3. Iterate all the links present in the graph. This is the primary service
   that the link graph component must provide to the other components
   (for example, the crawler and PageRank calculator) that comprise the
   Links 'R' Us project.
4. Insert an edge into the graph or refresh the UpdatedAt value of an
   existing edge.
5. Iterate the list of edges in the graph. This functionality is required
   by the PageRank calculator component.
6. Delete stale links that originated from a particular link and were not
   updated during the last crawler pass

