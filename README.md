# Mongo Query

`mongoquery` is a Go package for parsing input strings into MongoDB queries with support for complex conditions.

 ![screenshot](https://raw.githubusercontent.com/Kryptonux/MongoQuery/main/images/example.png)

## Installation

To use `mongoquery`, you can install it with `go get`:

```bash
go get github.com/Kryptonux/MongoQuery
```

## Example
Here's an example of how to use mongoquery to construct and execute MongoDB queries:
```go
package main

import (
	"fmt"

	"github.com/Kryptonux/MongoQuery/mongoquery"
)

func main() {
	// MongoDB connection details
	mongoURI := "mongodb://localhost:27017"
	databaseName := "your-database"
	collectionName := "your-collection"

	// Create a MongoDB client
	client, err := mongoquery.New(mongoURI, databaseName, collectionName)
	if err != nil {
		fmt.Printf("Error creating MongoDB client: %v\n", err)
		return
	}
	defer client.Close()

	// Input query string
	inputString := "_login.email_address: meow@protonmail.ch + _login.password_cracked: *"

	fmt.Println("Input Query: ", inputString)
	fmt.Println("MQuery Out.: ", mongoquery.ParseInputString(inputString))
	fmt.Println()

	// Execute the MongoDB query
	results, err := client.Query(inputString)
	if err != nil {
		fmt.Printf("Error executing MongoDB query: %v\n", err)
		return
	}
  
  // Process documents how ever you want but in this example we'll output how many results there is
	fmt.Printf("Matching documents: %v\n", len(results))
}
```
