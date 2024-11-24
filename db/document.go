package db

import (
	"encoding/json"
	"fmt"
	"os"
)

// Document represents a JSON document stored in the database.
type Document struct {
	ID   string                 `json:"id"`
	Data map[string]interface{} `json:"data"`
}

// SaveDocument saves a Document to a file.
func SaveDocument(path string, doc Document) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(doc)
}

// LoadDocument loads a Document from a file.
func LoadDocument(path string) (Document, error) {
	var doc Document
	file, err := os.Open(path)
	if err != nil {
		return doc, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&doc)
	if err != nil {
		return doc, err
	}

	return doc, nil
}

