package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/rhinoman/couchdb-go"
)

var database string

// Payload is example code
type Payload struct {
	Stuff Data
}

// Data is example code
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

// Fruits is example code
type Fruits map[string]int

// Vegetables is example code
type Vegetables map[string]int

func main() {
	if len(os.Args) == 1 { /* os.Args[0] is "compiler" or "compiler.exe" */
		fmt.Println("P-Put it i-in me, Senpai!")
		//return
	} else if len(os.Args) == 2 {
		filename := os.Args[1]
		fmt.Println("You put it in me, Senpai! <3", filename)
	} else if len(os.Args) > 2 { /* os.Args[0] is "main" or "main.exe" */
		fmt.Println("But Senpai, I can't Handle so many!")
		return
	}
	type TestDocument struct {
		Title string
		Note  string
	}
	type MyTest struct {
		Title string
	}

	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, timeout)
	if err != nil {
		fmt.Println("error")
	}
	auth := couchdb.BasicAuth{Username: "admin", Password: "password"}
	dbname := "mydatabase"
	created := conn.CreateDB(dbname, &auth)
	//*
	if created != nil {
		fmt.Println(created)
	}
	//*/
	db := conn.SelectDB(dbname, &auth)
	dblist, err := conn.GetDBList()
	fmt.Println(dblist)
	if dblist[2] == dbname {
		fmt.Println("yeah")
	}
	theDoc := TestDocument{
		Title: "My Document",
		Note:  "This is a note",
	}

	theDoc2 := TestDocument{
		Title: "",
		Note:  "",
	}
	theTest := MyTest{
		Title: "My Document",
	}
	theID := genUUID() //use whatever method you like to generate a uuid
	//The third argument here would be a revision, if you were updating an existing document
	blibb, err := db.Read(theID, theDoc2, nil)
	fmt.Println("before")
	fmt.Println(blibb)
	fmt.Println(theDoc.Title)
	fmt.Println(theDoc.Note)
	fmt.Println(theDoc)
	fmt.Println("after")
	blibb, err = db.Read(theID, theTest, nil)
	fmt.Println("before")
	fmt.Println(blibb)
	fmt.Println(theTest.Title)
	//fmt.Println(theTest.Note) struct doesn't have a Note Field
	fmt.Println(theTest)
	fmt.Println("after")

	rev, err := db.Save(theDoc, theID, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rev)
	//If all is well, rev should contain the revision of the newly created
	//or updated Document
	bla := "asd"
	if bla == "asdd" {
		//http section
		http.HandleFunc("/", serveRest)
		http.ListenAndServe("localhost:1337", nil)
	}

}

func genUUID() string {
	return "sddddad2"
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.URL)
	response, err := getJSONRespnse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func getJSONRespnse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apple"] = 25
	fruits["Orange"] = 11
	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", "  ")
}

/*

func IsValidCategory(category string) bool {
    switch category {
    case
        "auto",
        "news",
        "sport",
        "music":
        return true
    }
    return false
}


visitedURL := map[string]bool {
    "http://www.google.com": true,
    "https://paypal.com": true,
}
if visitedURL[thisSite] {
    fmt.Println("Already been here.")
}
*/
