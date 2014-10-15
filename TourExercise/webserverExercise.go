package main

import (
    "fmt"
    "net/http"
)

type String string

type Struct struct {
	Greeting string
	Punct string
	Who string
}

func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
		fmt.Fprint(w, str)
}

func (stct Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){

		msg := stct.Greeting +stct.Punct+stct.Who
		fmt.Fprint(w, msg)
}

func main() {
    http.Handle("/disha", String("Disha is genius"))
    http.Handle("/guna", &Struct{"Gunapal",":", "Awesome!"})
    http.ListenAndServe("localhost:4000", nil)
}