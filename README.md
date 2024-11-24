# Go-based NoSQL Database (MongoDB-like)

This project is a simple, in-memory NoSQL database system built using **Go (Golang)**. The database mimics the core functionality of **MongoDB**, where data is stored as **JSON documents** in **collections**, and the collections are organized into **folders**. The database supports concurrent operations using Go's goroutines and channels, and file-based persistence is used to save and load documents.

## Features

- **Collection-based Structure**: Documents are stored in collections, and each collection is represented by a folder.
- **JSON Document Storage**: Each document is stored as a JSON file.
- **Concurrency**: Uses **goroutines**, **channels**, and **mutexes** to handle concurrent operations safely.
- **Database and Collection Management**: Allows for the creation of databases and collections, and the addition and retrieval of documents.
- **File Persistence**: Documents are saved to files for persistence and can be loaded from disk.


### Description:

1. **`cmd/main.go`**: This is the entry point of the application where the program starts. It initializes a new database, adds collections, and performs actions like adding and retrieving documents.

2. **`db/database.go`**: Contains the logic for managing databases. It allows you to create a new database, add collections, and retrieve collections.

3. **`db/collection.go`**: Defines the operations for managing collections within a database. It allows you to add documents and retrieve them by their ID.

4. **`db/document.go`**: Handles reading and writing documents to disk. Each document is stored as a JSON file in the file system.

5. **`db/concurrency.go`**: (Optional) Implements concurrency features like adding documents asynchronously using goroutines and channels.

6. **`db/utils.go`**: (Optional) Includes helper functions that are used across the project, like printing database information or logging.

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yourusername/mongodb-like-go.git
   cd mongodb-like-go

2. Install the Go dependencies (if needed):
   
   ```bash
   go mod tidy

3. Build the project:
   
   ```bash
    go build

4. Run the application:

    ```bash
    go run cmd/main.go
    ```

## To do
- **Indexing**: Implement indexing for faster lookups based on document fields.
- **Querying**: Add query support to filter and retrieve documents based on specific conditions.
- **Authentication & Authorization**: Implement user roles and permissions for access control.
