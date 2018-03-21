// file mux_ex.go
package main

/*
NOTE: Ideally templates should be rendered with Angular/React or something similar
However, the coding challenge was to test golang skills, so rendering templates with golang itself.
*/

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	// Read json content
	fileContents, _ := ioutil.ReadFile("./json/response.json")
	jsonContent := bytes.NewReader(fileContents)

	// Unmarshal json
	var data interface{}
	if err := json.NewDecoder(jsonContent).Decode(&data); err != nil {
		log.Println("Invalid json ", err)
	}

	// Render data in template
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, data)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", index)

	log.Println("Starting server on localhost:80")
	log.Fatal(http.ListenAndServe(":80", r))
}
