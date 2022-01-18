package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hello World 2`))
	}).Methods(http.MethodGet)
	println("listen to 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err)
	}
}
