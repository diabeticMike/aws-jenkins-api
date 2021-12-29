package main

import (
	"github.com/gorilla/mux"
	"net/http"
)


func main(){
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello World`))
	}).Methods(http.MethodGet)

	if err := http.ListenAndServe(":8080", router); err != nil{
		panic(err)
	}
}