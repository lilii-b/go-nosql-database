package db

// AddDocumentAsync adds a document to the collection asynchronously.
func (db *Database) AddDocumentAsync(collectionName string, doc Document, ch chan error) {
	go func() {
		collection, err := db.GetCollection(collectionName)
		if err != nil {
			ch <- err
			return
		}
		err = collection.AddDocument(doc)
		ch <- err
	}()
}

