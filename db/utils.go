package db

import "fmt"

func PrintDatabaseInfo(db *Database) {
	fmt.Printf("Database: %s\n", db.Name)
	for _, collection := range db.Collections {
		fmt.Printf("Collection: %s\n", collection.Name)
	}
}

