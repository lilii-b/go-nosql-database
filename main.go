package main

import (
	"fmt"
	"log"
	"mydb/db"
)

func main() {
	// Create a new database
	database := db.NewDatabase("mydb")

	// Add a collection to the database
	usersCollection := database.AddCollection("users")

	// Create a new document
	doc := db.Document{
		ID:   "1",
		Data: map[string]interface{}{"name": "Alice", "age": 30},
	}

	// Add document asynchronously
	ch := make(chan error)
	database.AddDocumentAsync("users", doc, ch)

	// Wait for the result
	err := <-ch
	if err != nil {
		log.Fatalf("Error adding document: %v", err)
	} else {
		fmt.Println("Document added successfully.")
	}

	// Retrieve the document
	retrievedDoc, err := usersCollection.GetDocument("1")
	if err != nil {
		log.Fatalf("Error retrieving document: %v", err)
	} else {
		fmt.Printf("Retrieved Document: %+v\n", retrievedDoc)
	}
}

