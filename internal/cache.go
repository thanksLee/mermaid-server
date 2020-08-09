package internal

import "fmt"

// DiagramCache provides the ability to cache diagram results.
type DiagramCache interface {
	// Store stores a diagram in the cache.
	Store(diagram *Diagram) error
	// Has returns true if we have a cache stored for the given diagram description.
	Has(diagram *Diagram) (bool, error)
	// Get returns a cached version of the given diagram description.
	Get(diagram *Diagram) (*Diagram, error)
	// GetAll returns all of the cached diagrams.
	GetAll() ([]*Diagram, error)
	// Delete deletes a cached version of the given diagram.
	Delete(diagram *Diagram) error
}

// NewDiagramCache returns an implementation of DiagramCache.
func NewDiagramCache() DiagramCache {
	return &inMemoryDiagramCache{
		idToDiagram: map[string]*Diagram{},
	}
}

// inMemoryDiagramCache is an in-memory implementation of DiagramCache.
type inMemoryDiagramCache struct {
	idToDiagram map[string]*Diagram
}

// Store stores a diagram in the cache.
func (c *inMemoryDiagramCache) Store(diagram *Diagram) error {
	id, err := diagram.ID()
	if err != nil {
		return fmt.Errorf("cannot get diagram ID: %w", err)
	}
	c.idToDiagram[id] = diagram
	return nil
}

// Has returns true if we have a cache stored for the given diagram description.
func (c *inMemoryDiagramCache) Has(diagram *Diagram) (bool, error) {
	id, err := diagram.ID()
	if err != nil {
		return false, fmt.Errorf("cannot get diagram ID: %w", err)
	}
	if d, ok := c.idToDiagram[id]; ok && d != nil {
		return true, nil
	}
	return false, nil
}

// Get returns a cached version of the given diagram description.
func (c *inMemoryDiagramCache) Get(diagram *Diagram) (*Diagram, error) {
	id, err := diagram.ID()
	if err != nil {
		return nil, err
	}
	if d, ok := c.idToDiagram[id]; ok && d != nil {
		return d, nil
	}
	return nil, nil
}

// GetAll returns all of the cached diagrams.
func (c *inMemoryDiagramCache) GetAll() ([]*Diagram, error) {
	res := make([]*Diagram, len(c.idToDiagram))
	i := 0
	for _, diagram := range c.idToDiagram {
		res[i] = diagram
		i++
	}
	return res, nil
}

// Delete deletes a cached version of the given diagram.
func (c *inMemoryDiagramCache) Delete(diagram *Diagram) error {
	id, err := diagram.ID()
	if err != nil {
		return err
	}
	if _, ok := c.idToDiagram[id]; ok {
		delete(c.idToDiagram, id)
		return nil
	}
	return nil
}
