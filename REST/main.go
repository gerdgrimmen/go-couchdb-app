package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	fmt.Println(r.URL)
	response, err := getJsonRespnse()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(response))
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("127.0.0.1:1337", nil)

}

func getJsonRespnse() ([]byte, error) {
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
