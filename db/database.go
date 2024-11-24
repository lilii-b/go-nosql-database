package db

import (
	"fmt"
	"sync"
)

// Database represents a collection of collections.
type Database struct {
	Name        string
	Collections map[string]*Collection
	mutex       sync.RWMutex
}

// NewDatabase creates a new empty database.
func NewDatabase(name string) *Database {
	return &Database{
		Name:        name,
		Collections: make(map[string]*Collection),
		mutex:       sync.RWMutex{},
	}
}

// AddCollection adds a new collection to the database.
func (db *Database) AddCollection(name string) *Collection {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	collection := NewCollection(name)
	db.Collections[name] = collection
	return collection
}

// GetCollection retrieves a collection by its name.
func (db *Database) GetCollection(name string) (*Collection, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	collection, exists := db.Collections[name]
	if !exists {
		return nil, fmt.Errorf("collection %s not found", name)
	}
	return collection, nil
}

