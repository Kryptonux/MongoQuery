package main

import (
	"fmt"

	"xks/mquery/mongoquery"
)

func main() {
	mongoURI := "mongodb://localhost:27017"
	databaseName := "xkeyscore"
	collectionName := "data"

	client, err := mongoquery.New(mongoURI, databaseName, collectionName)
	if err != nil {
		fmt.Printf("Error creating MongoDB client: %v\n", err)
		return
	}
	defer client.Close()

	inputString := "_login.email_address: meow@protonmail.ch + _login.password_cracked: *"

	fmt.Println("\nInput Query: ", inputString)
	fmt.Println("MQuery Out.: ", mongoquery.ParseInputString(inputString))
	fmt.Println()

	results, err := client.Query(inputString)
	if err != nil {
		fmt.Printf("Error executing MongoDB query: %v\n", err)
		return
	}

	fmt.Printf("Matching documents: %v", len(results))
}
