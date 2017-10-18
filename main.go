package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rhinoman/couchdb-go"
)

var database string

// Payload is example code
/*type Payload struct {
	Stuff Data
}
//*/

// Data is example code
/*type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}
//*/

// MyTest is
type MyTest struct {
	Title string
}

// Fruits is example code
//type Fruits map[string]int

// Vegetables is example code
//type Vegetables map[string]int

func main() {
	/*
		if len(os.Args) == 1 { // os.Args[0] is "compiler" or "compiler.exe"
			fmt.Println("P-Put it i-in me, Senpai!")
			//return
		} else if len(os.Args) == 2 {
			filename := os.Args[1]
			fmt.Println("You put it in me, Senpai! <3", filename)
		} else if len(os.Args) > 2 { // os.Args[0] is "main" or "main.exe"
			fmt.Println("But Senpai, I can't Handle so many!")
			return
		}
	//*/
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
	created := conn.CreateDB("mydatabase", &auth)
	if created != nil {
		return
	}
	db := conn.SelectDB("mydatabase", &auth)
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

	theID := genUUID()
	blibb, err := db.Read(theID, theDoc2, nil)
	fmt.Println(blibb)
	/*
		fmt.Println(theDoc.Title)
		fmt.Println(theDoc.Note)
		fmt.Println(theDoc)
		//*/
	blibb, err = db.Read(theID, theTest, nil)
	/*
		fmt.Println(theTest.Title)
		//fmt.Println(theTest.Note)
		fmt.Println(theTest)
		//*/
	rev, err := db.Save(theDoc, theID, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rev)
	bla := "asd"
	if bla == "asdd" {
		//http section
		http.HandleFunc("/mydatabase", serveRestCouchDB)
		//http.HandleFunc("/", serveRest)
		http.ListenAndServe("localhost:1337", nil)
	}

}

func genUUID() string {
	return "sad2"
}

/*
func serveRest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.URL)
	response, err := getJSONRespnse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}
//*/
func serveRestCouchDB(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.URL)
	response, err := simpleSelect()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}
func simpleDBCreate() (string, error) {
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, 5000)
	if err != nil {
		fmt.Println("error")
	}
	auth := couchdb.BasicAuth{Username: "admin", Password: "password"}
	created := conn.CreateDB("mydatabase", &auth)
	if created != nil {
		return "", err
	}
	return "{ok}.", nil
}
func simpleSelect() (string, error) {
	conn, err := couchdb.NewConnection("127.0.0.1", 5984, 5000)
	if err != nil {
		fmt.Println("error")
	}
	auth := couchdb.BasicAuth{Username: "admin", Password: "password"}
	//created := conn.CreateDB("mydatabase", &auth)
	/*if created != nil {
		fmt.Println(created)
	}*/
	db := conn.SelectDB("mydatabase", &auth)
	myt := MyTest{
		Title: "asd",
	}
	blibb, err := db.Read("sad", myt, nil)
	if err != nil {
		return "", err
	}
	return blibb, nil
}

/*
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

//*/
