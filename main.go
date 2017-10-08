package main

import (
	"fmt"
	"time"

	"github.com/rhinoman/couchdb-go"
)

func main() {
	fmt.Println("asd")
	type TestDocument struct {
		Title string
		Note  string
	}

	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	if err != nil {
		fmt.Println("error")
	}
	auth := couchdb.BasicAuth{Username: "admin", Password: "password"}
	created := conn.CreateDB("mydatabase", &auth)
	if created != nil {
		fmt.Println(created)
	}
	db := conn.SelectDB("mydatabase", &auth)
	theDoc := TestDocument{
		Title: "My Document",
		Note:  "This is a note",
	}

	theID := genUUID() //use whatever method you like to generate a uuid
	//The third argument here would be a revision, if you were updating an existing document
	rev, err := db.Save(theDoc, theID, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rev)
	//If all is well, rev should contain the revision of the newly created
	//or updated Document

}

func genUUID() string {
	return "sad"
}
