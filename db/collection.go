package db

import (
	"fmt"
	"sync"
)

// Collection represents a group of documents.
type Collection struct {
	Name    string
	Docs    map[string]Document
	mutex   sync.RWMutex
}

// NewCollection creates and returns a new collection.
func NewCollection(name string) *Collection {
	return &Collection{
		Name:    name,
		Docs:    make(map[string]Document),
		mutex:   sync.RWMutex{},
	}
}

// AddDocument adds a document to the collection.
func (c *Collection) AddDocument(doc Document) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.Docs[doc.ID] = doc
	return SaveDocument(fmt.Sprintf("./mydb/%s/%s/%s.json", c.Name, doc.ID, doc.ID), doc)
}

// GetDocument retrieves a document by its ID.
func (c *Collection) GetDocument(id string) (Document, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	doc, exists := c.Docs[id]
	if !exists {
		return Document{}, fmt.Errorf("document with ID %s not found", id)
	}
	return doc, nil
}

